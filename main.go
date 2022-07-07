package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

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

	//file, error := os.Create("posts.csv")
	//CheckError(error)

	//writer := csv.NewWriter(file)

	var posts []Post

	doc.Find("div.ecs-posts").Find("div.elementor-section-wrap").Each(func(index int, item *goquery.Selection) {
		h3 := item.Find("h3")

		title := h3.Text()
		url, _ := h3.Find("a").Attr("href")
		term := item.Find("div.elementor-text-editor").Text()

		posts = append(posts, Post{
			Title: title,
			Url:   url,
			Term:  term,
		})
	})

	j, _ := json.Marshal(posts)

	log.Println(string(j))
}
