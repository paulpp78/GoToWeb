package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	println("Le serveur écoute sur le port :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}
