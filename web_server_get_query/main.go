package main

import (
	"fmt"
	"net/http"
	"web_server_get_query/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/articles", handler.ArticleHandler)
	mux.HandleFunc("/get-post", handler.PostGet)

	fmt.Print("Starting Serve on port 8000")
	err := http.ListenAndServe(":8000", mux)
	fmt.Print("Error ", err)
}
