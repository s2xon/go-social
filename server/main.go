package main

import (
	"fmt"
	"net/http"
	//"fmt"
	"log"
	"root/server/api/facebook"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fb.Login(), http.StatusSeeOther)
}

func AccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	user := fb.AccessToken(r)

	fmt.Println(user.Access_Token)
}

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/facebook_redirect", AccessTokenHandler)
	fmt.Println("Starting server @ https://localhost:8080")
	log.Fatal(http.ListenAndServeTLS(":8080", "../certificates/localhost.pem", "../certificates/localhost-key.pem", nil))

}
