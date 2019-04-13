package clck

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// ShortURL returns short URL from clck.ru service
func ShortURL(longURL string) (string, error) {
	client := &http.Client{}
	resp, err := client.Get(`https://clck.ru/--?url=` + url.QueryEscape(longURL))
	if err != nil {
		return longURL, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
