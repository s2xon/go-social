package main

import (
	"net/http"
	//"fmt"
	"log"
  "root/server/api/facebook"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, fb.Login(), http.StatusSeeOther)
}

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

  http.HandleFunc("/login", LoginHandler)
	log.Fatal(http.ListenAndServeTLS(":8080", "../certificates/localhost.pem", "../certificates/localhost-key.pem", nil))
}
