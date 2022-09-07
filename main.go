package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Article struct {
	Id       uint16
	Title    string
	Anons    string
	FullText string
}

// * масив постів з елементами типу Article
var posts = []Article{}
var showPost = Article{}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	// !	треба в параметрах передавати всі темплейти, які ми будемо підключати !

	// перевірка на помилку
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//  * підключаємось до бази даних, щоб отримати всі статті
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`")

	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)

		if err != nil {
			panic(err)
		}

		// fmt.Println(fmt.Sprintf("Post: %s with id %d", post.Title, post.Id))
		posts = append(posts, post)
	}

	//	* всередині HTML файла буде певний блок, який буде називатись index
	// всередині шаблона буде динамічне підключення
	// другий параметр показує який конкретно блок ми намагаємось вивести
	//	третій параметр -- настройки
	t.ExecuteTemplate(w, "index", posts)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)

}

func save_article(w http.ResponseWriter, r *http.Request) {
	//  * r -- вся інформація зі сторінки
	//  ? FormValue приймає назву елемента HTML
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	// ! валідація форми
	if title == "" || anons == "" || full_text == "" {
		// todo виводимо помилку на екран
		// fmt.Fprintf(w, "Не все данные заполнены")
		// todo робимо редірект на форму
		http.Redirect(w, r, "/create", http.StatusSeeOther)
	} else {

		//  ! підключаємось до БД
		db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/golang")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		// * додавання/установка даних
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))

		if err != nil {
			panic(err)
		}
		defer insert.Close()

		//  * робимо переадресацію після додавання в базу даних
		//  todo передаємо всі параметри w i r, сторінку, на яку буде переадресація і код відповіді
		http.Redirect(w, r, "/", 301)
	}
}

func show_post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // ~ & створюємо об'єкт з параметрами з динамічної адреси

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// ^ робимо запит в БД
	res, err := db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id`='%s'", vars["id"]))

	if err != nil {
		panic(err)
	}

	showPost = Article{}
	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)

		if err != nil {
			panic(err)
		}

		showPost = post
	}
	t.ExecuteTemplate(w, "show", showPost)
}

func handleFunc() {
	// * відслідковуємо url адреси за допомогою gorilla/mux
	// *створили новий об'єкт роутер
	rtr := mux.NewRouter()
	// ! обробка статичних файлів
	// обробляeмо всі url адреси, які починаються з static
	// * кожного разу, коли буде йти звернення до static, ми з цього звернення видаляємо слово static
	// * а далі по шляху, який залишається, ми шукаємо потрібний файл в папці static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create", create).Methods("GET")
	rtr.HandleFunc("/save_article", save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

	http.Handle("/", rtr) // & обробка всіх url адрес буде відбуватись через router
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
