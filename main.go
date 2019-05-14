package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func indexHandler(writer http.ResponseWriter, request *http.Request)  {
	template, err := template.ParseFiles("templates/index.html", "templates/head.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(writer, err.Error())
	}

	template.ExecuteTemplate(writer, "index", nil)
}

func loginHandler(writer http.ResponseWriter, request *http.Request)  {
	template, err := template.ParseFiles("templates/login.html", "templates/head.html", "templates/footer.html")
	if err != nil {
		fmt.Fprint(writer, err.Error())
	}

	template.ExecuteTemplate(writer, "login", nil)
}

func main() {
	fmt.Println("Listening Port 8080")

	//
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login.html", loginHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

