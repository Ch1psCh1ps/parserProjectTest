package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func writeFile(data, filename string) {
	file, error := os.Create(filename)
	defer file.Close()
	checkError(error)

	file.WriteString(data)
}

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

	ecs, error := doc.Find("div.ecs-posts").Html()
	checkError(error)

	//fmt.Println(ecs)
	writeFile(ecs, "writeFile.html")
}
