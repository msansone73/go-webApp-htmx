package main

import (
	"main/rotes"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	rotes.SetRotes(mux)

	fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fileServer))


	http.ListenAndServe(":8080", mux)
}






