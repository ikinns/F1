package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var tld, strout string
var visited map[string]bool
var f1file *os.File

func processPage(s string) {
	visited[s] = true
	doc, err := goquery.NewDocument(s)
	if err != nil {
		fmt.Println("\nPage - ", s, " NOT FOUND")
		strout += "Page - " + s + " NOT FOUND\n"
	} else {

		//	fmt.Println("\nPage - ", s, visited[s])
		fmt.Println("\nPage - ", s)
		strout += "Page - " + s + "\n"
		time.Sleep(1 * time.Second)
		doc.Find(`table.wikitable`).Each(processContent)
/*
		doc.Find(`li a`).
			Each(func(i int, e *goquery.Selection) {
				link, _ := e.Attr("href")
				link = strings.TrimSpace(link)
				if link != "./" {
					if !strings.HasPrefix(link, "http") {
						link = tld + link
					}
					if !visited[link] {
						processPage(link)
					}
				}
			})
*/
	}
}

func processContent(i int, body *goquery.Selection) {
	// goquery.NodeName(e)
	strout += "\n"
	body.Find(`tr`).Each(func(i int, row *goquery.Selection) {
		rowtxt := ""
		row.Find(`th, td`).Each(func(i int, cell *goquery.Selection) {
			item := cell.Text()
			item = strings.TrimSpace(item)
			item = strings.ReplaceAll(item, "\t", "")
			item = strings.ReplaceAll(item, "\r", "")
			item = strings.ReplaceAll(item, "\n", "")
			rowtxt += item + ", "
		})
		strout += rowtxt + "\n"
	})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	visited = make(map[string]bool)
	for year := 1950; year < 2020; year++ {
		fname := "f1_" + strconv.Itoa(year) + ".csv"
		if year < 1981 {
			tld = "https://en.wikipedia.org/wiki/" + strconv.Itoa(year) + "_Formula_One_season"
		} else {
			tld = "https://en.wikipedia.org/wiki/" + strconv.Itoa(year) + "_Formula_One_World_Championship"
		}
		f1file, err := os.Create(fname)
		check(err)
		defer f1file.Close()

		strout = ""
		processPage(tld)
		f1file.WriteString(strout)
	}
}
