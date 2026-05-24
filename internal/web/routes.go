package web

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// pages
	mux.HandleFunc("/", basePage)

	// api
	mux.HandleFunc("POST /todos", basePage)

	// static
	fs := http.FileServer(http.Dir("../static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
