package main

import (
	"fmt"
)

//Это файл для помощи с поиском нужных строк на сайте и функции проверки ошибки

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

/*func writeFile(data, filename string) {
	file, error := os.Create(filename)
	defer file.Close()
	CheckError(error)

	file.WriteString(data)
}

func help(url string) {

	response, error := http.Get(url)
	defer response.Body.Close()
	CheckError(error)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, error := goquery.NewDocumentFromReader(response.Body)

	ecs, error := doc.Find("div.ecs-posts").Html()
	CheckError(error)

	writeFile(ecs, "writeFile.html")
}
*/
/*func main() {
	help("https://megasport.msk.ru/afisha-meropriyatiy/")
}*/
