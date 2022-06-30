// Шаблон для роботи с go та http запросами( веб сайти і тд)
package main

import (
	"log"
	"net/http"
	"www/server"
)

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

	//провірка директорії
	// dir, err := os.Getwd()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
}
func run() error {
	tmp := []string{"templates/index.html", "templates/header.html", "templates/footer.html"}
	start := server.Init()
	start.Handle("/static/")
	start.Handle("/content/")
	start.Request("index", "/", tmp)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
