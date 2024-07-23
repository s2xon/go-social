package main

import (
	"fmt"
	"net/http"
	//"fmt"
	"log"
	"root/server/api/facebook"
)

var User *fb.User

func Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static")

	if User != nil {
		fmt.Println("I STORE THE VARIABLE HERE", User)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fb.Login(), http.StatusSeeOther)
}

func AccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	User = fb.AccessToken(r)

	fmt.Println("here is the user atkn", User)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/facebook_redirect", AccessTokenHandler)
	fmt.Println("Starting server @ https://localhost:8080")
	log.Fatal(http.ListenAndServeTLS(":8080", "../certificates/localhost.pem", "../certificates/localhost-key.pem", nil))

}
