package main

import (
	"fmt"
	"net/http"
) // пакет для функцій по виводу інформації в консолі чи на сайті

func home_page(page http.ResponseWriter, r *http.Request) { // через параметр page ми зможемо звертатись до сторінки, показувати, виводити текст, HTML
	// r параметр, який передається, ми зможемо відслідкувати дані при запиті/підключенні
	fmt.Fprintf(page, "Go is super easy!") //ствроюємо форматовану строку, в яку можна динамічно підставляти значення
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
	handleRequest()
}
