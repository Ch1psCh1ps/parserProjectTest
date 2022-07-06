package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//fmt.Println("Hi. Thats enough for first commit (:")

	url := "https://megasport.msk.ru/afisha-meropriyatiy/"

	response, error := http.Get(url)
	defer response.Body.Close()
	checkError(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, error := goquery.NewDocumentFromReader(response.Body)
	checkError(error)

	elementor := doc.Find("div.elementor").Size()

	fmt.Println(elementor)
}
