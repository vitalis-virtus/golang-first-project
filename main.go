package main

import (
	"fmt"
	"html/template"
	"net/http"
)

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

func handleFunc() {
	// ! обробка статичних файлів
	// обробляeмо всі url адреси, які починаються з static
	// * кожного разу, коли буде йти звернення до static, ми з цього звернення видаляємо слово static
	// * а далі по шляху, який залишається, ми шукаємо потрібний файл в папці static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
