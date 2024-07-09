package main

import (
	"net/http"
	//"fmt"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServeTLS(":8080", "../certificates/localhost.pem", "../certificates/localhost-key.pem", nil))
}
