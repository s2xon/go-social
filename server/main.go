package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	//"fmt"
	"log"
	"root/server/api/facebook"
)

var User *fb.User

func Handler(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "../static/index.html")

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

func ImageHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("fileToUpload")
	if err != nil {
		log.Println("error retrieving file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fmt.Println("1")
	dst, err := os.Create(handler.Filename)
	if err != nil {
		log.Println("error creating file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	fmt.Println("1")

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "uploaded file")
}

func main() {

	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Handler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/facebook_redirect", AccessTokenHandler)

	http.HandleFunc("/upload_image", ImageHandler)

	fmt.Println("Starting server @ https://localhost:8080")
	log.Fatal(http.ListenAndServeTLS(":8080", "../certificates/localhost.pem", "../certificates/localhost-key.pem", nil))

}
