package main

import (
	"fmt"
	"net/http"
) // пакет для функцій по виводу інформації в консолі чи на сайті

// створюємо структуру
type User struct {
	name                  string  // відразу треба вказувати тип даних
	age                   uint16  //ціле невід'ємне число
	money                 int16   //ціле число
	avg_grades, happiness float64 //число з комою
}

// створюємо метод для структури
func (u User) getAllInfo() string { // u -- об'єкт до якого ми будемо звертатись в методі
	return fmt.Sprintf("Username is: %s. He is %d and he has money"+
		" equal: %d.", u.name, u.age, u.money) // %s-для строки	%d-для числа
}

func (u *User) setNewName(newName string) { // треба передавати посилання на сам об'єкт, викор знак "*"
	u.name = newName
}

func home_page(page http.ResponseWriter, r *http.Request) { // через параметр page ми зможемо звертатись до сторінки, показувати, виводити текст, HTML
	// r параметр, який передається, ми зможемо відслідкувати дані при запиті/підключенні
	tom := User{"Tom", 35, -10, 2.2, 0.6}
	tom.setNewName("Alex")
	fmt.Fprint(page, tom.getAllInfo()) //ствроюємо форматовану строку, в яку можна динамічно підставляти значення
	// fmt.Fprintf(page, tom.name)
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
