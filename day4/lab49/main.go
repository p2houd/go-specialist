package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, pagesChan chan Page) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	// работает по принципу lifo
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	pagesChan <- Page{
		URL:  url,
		Size: len(body),
	}
	fmt.Println(len(body))
}

func main() {
	pagesChan := make(chan Page)

	pages := []string{
		"https://github.com",
		"https://honda.cn",
		"http://honda.cn",
		"https://yandex.ru",
	}

	for _, page := range pages {
		go responseSize(page, pagesChan)
	}
	for i := 0; i < len(pages); i++ {
		page := <-pagesChan
		fmt.Println(page.Size, page.URL)
	}
}
