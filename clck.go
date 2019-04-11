package clck

import (
	"io/ioutil"
	"net/http"
)

// ShortURL returns short url from clck.ru service
func ShortURL(url string) (string, error) {
	client := &http.Client{}
	resp, err := client.Get(`https://clck.ru/--?url=` + url)
	if err != nil {
		return url, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}
