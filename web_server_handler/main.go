package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//Cara kedua membuat handler memasukan function kedalam variabel
	loginHandler := func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)

		w.Write([]byte("Ini Login Page"))
	}

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/login", loginHandler)

	//Cara ketiga membuat handler (dengan anonymous function)
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ini SignUp Page"))
	})

	log.Print("Starting webserver in port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal((err))
}

//cara pertama membuat handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Ini HomePage"))
}
