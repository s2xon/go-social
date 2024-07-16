package fb

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func Login() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	u, err := url.Parse("https://www.facebook.com/v20.0/dialog/oauth?client_id={app-id}&redirect_uri={redirect-uri}&state={state-param}&config_id={config}")
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("client_id", os.Getenv("FB_ID"))
	q.Set("redirect_uri", os.Getenv("MY_URI"))
	q.Set("config_id", os.Getenv("FB_CONFIG"))
	u.RawQuery = q.Encode()

	return u.String()
}

func AccessToken(r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	u, err := url.Parse(r.URL.RequestURI())
	if err != nil {
		log.Fatal(err)
	}

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
	fmt.Println(resp)

}

// func Upload() {
//
//}
