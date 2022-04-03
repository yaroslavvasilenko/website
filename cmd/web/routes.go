package main

import "net/http"

//  распределение http запроса
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/note", app.showNote)
	mux.HandleFunc("/note/create", app.createNote)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
