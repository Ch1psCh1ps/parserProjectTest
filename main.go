package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Post struct {
	titles string `json:"titles"`
	urls   string `json:"urls"`
	term   string `json:"term"`
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

	doc.Find("div.ecs-posts").Find("div.elementor-section-wrap").Each(func(index int, item *goquery.Selection) {
		h3 := item.Find("h3")

		title := h3.Text()
		url, _ := h3.Find("a").Attr("href")
		term := item.Find("div.elementor-text-editor").Text()

		posts := []Post{
			{
				titles: title,
				urls:   url,
				term:   term,
			},
		}

		fmt.Println(posts)

		//j, _ := json.Marshal(posts)

		//log.Println(string(j))
		//writer.Write(posts)
	})

	//writer.Flush()

}
