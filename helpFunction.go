package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"parserProjectTest/lib"
)

//Это файл для помощи с поиском нужных строк на сайте

func writeFile(data, filename string) {
	file, error := os.Create(filename)
	defer file.Close()
	lib.CheckError(error)

	file.WriteString(data)
}

func help(url string) {

	response, error := http.Get(url)
	defer response.Body.Close()
	lib.CheckError(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, error := goquery.NewDocumentFromReader(response.Body)

	ecs, error := doc.Find("div.ecs-posts").Html()
	lib.CheckError(error)

	writeFile(ecs, "writeFile.html")
}

/*func main() {
	help("https://megasport.msk.ru/afisha-meropriyatiy/")
}*/
