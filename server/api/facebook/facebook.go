package fb

import (
	"log"
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
