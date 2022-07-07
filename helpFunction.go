package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
)

/*Это файл для помощи с поиском нужных строк на сайте и отдельной функции по поиску ошибок*/
func writeFile(data, filename string) {
	file, error := os.Create(filename)
	defer file.Close()
	checkError(error)

	file.WriteString(data)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
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

	ecs, error := doc.Find("div.ecs-posts").Html()
	checkError(error)

	writeFile(ecs, "writeFile.html")
}
