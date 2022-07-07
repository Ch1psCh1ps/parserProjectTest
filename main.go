package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func main() {

	url := "https://megasport.msk.ru/afisha-meropriyatiy/"

	response, error := http.Get(url)
	defer response.Body.Close()
	checkError(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, error := goquery.NewDocumentFromReader(response.Body)
	checkError(error)

	file, error := os.Create("posts.csv")
	checkError(error)

	writer := csv.NewWriter(file)

	doc.Find("div.ecs-posts").Find("div.elementor-section-wrap").Each(func(index int, item *goquery.Selection) {
		h3 := item.Find("h3")
		title := h3.Text()
		url, _ := h3.Find("a").Attr("href")

		time := item.Find("div.elementor-text-editor").Text()

		posts := []string{title, url, time}

		writer.Write(posts)
	})

	writer.Flush()

}
