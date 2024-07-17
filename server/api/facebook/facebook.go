package fb

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type User struct {
	Access_Token string `json:"access_token"`
}

func Login() string {
	fmt.Println("-----------------loginhere--------------------")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	u, err := url.Parse("https://www.facebook.com/v20.0/dialog/oauth?client_id={app-id}&redirect_uri={redirect-uri}&state={state-param}&config_id={config}")
	if err != nil {
		log.Fatal(err)
	}

	ranNum := rand.Intn(10)
	q := u.Query()
	q.Set("client_id", os.Getenv("FB_ID"))
	q.Set("redirect_uri", os.Getenv("MY_URI"))
	q.Set("state", strconv.Itoa(ranNum))
	q.Set("config_id", os.Getenv("FB_CONFIG"))

	u.RawQuery = q.Encode()

	fmt.Println(u.String())

	return u.String()
}

func AccessToken(r *http.Request) *User {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	u, err := url.Parse(r.URL.RequestURI())
	if err != nil {
		log.Fatal(err)
	}
  fmt.Println(u)
	m, _ := url.ParseQuery(u.RawQuery)
	code := m["code"][0]

	s, err := url.Parse("https://graph.facebook.com/v20.0/oauth/access_token?client_id={app-id}&redirect_uri={redirect-uri}&client_secret={app-secret}&code={code-parameter}")
	if err != nil {
		panic(err)
	}

	q := s.Query()
	q.Set("client_id", os.Getenv("FB_ID"))
	q.Set("redirect_uri", os.Getenv("MY_URI"))
	q.Set("client_secret", os.Getenv("FB_SECRET"))
	q.Set("code", code)
	s.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", s.String(), nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf := new(strings.Builder)
	n, err := io.Copy(buf, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(buf.String())

	var user *User
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, user)
	if err != nil {
		panic(err)
	}
	return user
}

// func Upload() {
//
//}
