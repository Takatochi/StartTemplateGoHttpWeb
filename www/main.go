// Шаблон для роботи с go та http запросами( веб сайти і тд)
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//старт темлейтов для метода index (головна сторінка)
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func hadleRequst() {
	//обробка шляху до стилів
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
func main() {

	hadleRequst()

	//провірка директорії
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
}
