package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Article Page ", pageNumber)
}

func PostGet(w http.ResponseWriter, r *http.Request) {

	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini Post"))
	case "POST":
		w.Write([]byte("Ini GET"))
	case "UPDATE":
		w.Write([]byte("Ini UPDATE"))
	default:
		http.Error(w, "Error", http.StatusBadRequest)
	}
}

