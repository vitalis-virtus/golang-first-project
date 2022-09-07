package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"` // відображаємо в джсон форматі
	Age  uint16 `json:"age"`
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	// !	треба в параметрах передавати всі темплейти, які ми будемо підключати !

	// перевірка на помилку
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	//	* всередині HTML файла буде певний блок, який буде називатись index
	// всередині шаблона буде динамічне підключення
	// другий параметр показує який конкретно блок ми намагаємось вивести
	//	третій параметр -- настройки
	t.ExecuteTemplate(w, "index", nil)
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
	// title := r.FormValue("title")
	// anons := r.FormValue("anons")
	// full_text := r.FormValue("full_text")

}

func handleFunc() {
	// ! обробка статичних файлів
	// обробляeмо всі url адреси, які починаються з static
	// * кожного разу, коли буде йти звернення до static, ми з цього звернення видаляємо слово static
	// * а далі по шляху, який залишається, ми шукаємо потрібний файл в папці static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)

	http.ListenAndServe(":8080", nil)
}

func main() {
	// handleFunc()
	// * в db записується саме підключення до БД
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// ! додавання/установка даних
	insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Bob', 35)")

	if err != nil {
		panic(err)
	}
	defer insert.Close()

	//  ! вибірка даних
	res, err := db.Query("SELECT `name`, `age` FROM `users`")
	if err != nil {
		panic(err)
	}

	// ? запустимо цикл по всіх результатах з сервера
	// * Next() видає чи є настпна строка, яку можна обробити
	for res.Next() {
		var user User
		// * Scan() перевіряє чи існує якесь певне значення
		//  * перебираємо кожен ряд і витягуємо два параметри Name i Age
		err = res.Scan(&user.Name, &user.Age)

		if err != nil {
			panic(err)
		}

		//  * Sprintf формує саму строку
		fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))

	}

}
