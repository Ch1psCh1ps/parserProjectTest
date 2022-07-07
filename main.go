package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type Post struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Term  string `json:"term"`
}

func main() {

	url := "https://megasport.msk.ru/afisha-meropriyatiy/"

	response, error := http.Get(url)
	defer response.Body.Close()
	CheckError(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, error := goquery.NewDocumentFromReader(response.Body)
	CheckError(error)

	var posts []Post

	doc.Find("div.ecs-posts").Find("div.elementor-section-wrap").Each(func(index int, item *goquery.Selection) {
		h3 := item.Find("h3")

		title := h3.Text()
		url, _ := h3.Find("a").Attr("href")
		term := item.Find("div.elementor-text-editor").Text()

		posts = append(posts, Post{
			Title: strings.TrimSpace(title),
			Url:   strings.TrimSpace(url),
			Term:  strings.TrimSpace(term),
		})
	})

	j, _ := json.Marshal(posts)

	fmt.Println(string(j))
}
