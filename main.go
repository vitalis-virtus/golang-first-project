package main

import (
	"fmt"           // пакет для функцій по виводу інформації в консолі чи на сайті
	"html/template" //пакет для вивод шабонів HTML
	"net/http"      // пакет для запуска сервера
)

// створюємо структуру
type User struct {
	Name                  string   // відразу треба вказувати тип даних
	Age                   uint16   //ціле невід'ємне число
	Money                 int16    //ціле число
	Avg_grades, Happiness float64  //число з комою
	Hobbies               []string //створюємо список строк
}

// створюємо метод для структури
func (u User) getAllInfo() string { // u -- об'єкт до якого ми будемо звертатись в методі
	return fmt.Sprintf("Username is: %s. He is %d and he has money"+
		" equal: %d.", u.Name, u.Age, u.Money) // %s-для строки	%d-для числа
}

func (u *User) setNewName(newName string) { // треба передавати посилання на сам об'єкт, викор знак "*"
	u.Name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) { // через параметр page ми зможемо звертатись до сторінки, показувати, виводити текст, HTML
	// r параметр, який передається, ми зможемо відслідкувати дані при запиті/підключенні
	tom := User{"Tom", 35, -10, 2.2, 0.6,
		[]string{"fotball", "dance", "skate"}}
	// tom.setNewName("Alex")
	// fmt.Fprint(page, tom.getAllInfo()) //ствроюємо форматовану строку, в яку можна динамічно підставляти значення
	// fmt.Fprintf(page, tom.name)

	// виводимо HTML на сторінці
	// fmt.Fprintf(page, `<h1>Header</h1>
	// <b>Text</b>`)

	//err змінна, яка буде зберігати якісь помилки
	// tmpl, err
	tmpl, _ := template.ParseFiles("templates/home_page.html") //ми підгружаємо певний HTML шаблон
	tmpl.Execute(page, tom)                                    //виконуємо шабон і передаємо в нього дані у вигляді об'єкта tom
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page) //роутинг -- при переході на "/" ми застосовуємо метод "home_page"
	http.HandleFunc("/contacts", contacts_page)
	http.ListenAndServe(":8080", nil) // запуск локального сервера на порті 8080 / nil - аналог null / дргуий параметр -- настройки
}

func main() {
	// var bob User = ...

	//наступні методи створення екземпляра структури однакові -- можна з ключамми, можна і без
	// bob := User{name: "Bob", age: 25, money: -50, avg_grades: 4.2, happiness: 0.8}

	handleRequest()
}
