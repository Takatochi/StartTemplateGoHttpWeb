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
	// hadleRequst()
	// server.Print()
	start := server.Init("/static/", "/")

	start.Handle()
	start.Request()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
