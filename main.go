package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	siteUrl := "https://oxu.az"
	res, err := http.Get(siteUrl)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		checkError(err)
	}(res.Body)

	checkError(err)

	if res.StatusCode > 400 {
		fmt.Println("Status Code:", res.StatusCode)
	}
	file, err := os.Create("oxuazTitle.csv")
	checkError(err)
	csvFileWriter := csv.NewWriter(file)

	reader, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	reader.Find("section.news-list").Find("div.news-i").Find("a.news-i-inner").Each(func(index int, item *goquery.Selection) {
		divTag := item.Find("div.title")
		title := strings.TrimSpace(divTag.Text())

		//// String slice to store all data
		postDetail := []string{title}
		//// To write csv File
		err := csvFileWriter.Write(postDetail)
		checkError(err)
	})

	csvFileWriter.Flush()
}
