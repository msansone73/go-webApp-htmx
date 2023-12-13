package main

import (
	"main/rotes"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	rotes.SetRotes(mux)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", mux)
}






