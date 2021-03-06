package useragent

import (
	"log"
	"testing"
)

func TestUserAgent(t *testing.T) {
	var ua UserAgent

	if ua.Random() == "" {
		t.Error("browser.Random is empty")
	}

	log.Println(ua.Random())
	log.Println(ua.Android())
	log.Println(ua.Chrome())
}
