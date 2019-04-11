package clck

import (
	"testing"
)

func TestShortURL(t *testing.T) {
	url, err := ShortURL("https://yandex.ru")
	if err == nil {
		t.Logf("Short URL received: %s", url)
	} else {
		t.Errorf("Got error %+v", err)
	}
}
