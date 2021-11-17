package dadjoke

import (
	"io"
	"log"
	"net/http"
	"time"
)

// Fetch returns a dad joke
func Fetch() string {
	url := "https://icanhazdadjoke.com/"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "text/plain")
	client := &http.Client{Timeout: time.Second * 5}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
