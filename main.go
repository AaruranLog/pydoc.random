package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	"github.com/gocolly/colly"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func writeDocToFile(filename, body string) {
	if _, err := os.Stat("docs/" + filename); err == nil {
		fmt.Println("File already written.")
		return
	}
	f, err := os.Create("docs/" + filename)
	check(err)
	defer f.Close()
	cleanedBody := strings.TrimSpace(body)
	cleanedBody = strings.Replace(cleanedBody, "¶", "\n")
	_, err = f.WriteString(cleanedBody)
	f.Sync()
	fmt.Println("Written to file:", filename)
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("docs.python.org", "http://docs.python.org",
		"https://docs.python.org", "https://docs.python.org/3/"),
	)

	c.OnResponse(func(r * colly.Response){
		fmt.Println("Response received from:", r.Request.URL)
	})

	// Basic error-handling
	c.OnError(func(_ *colly.Response, err error){
		fmt.Println("Something went wrong during request:", err)
		check(err)
	})

	c.OnHTML("div[class=document]", func(e *colly.HTMLElement) {
		documentBody := e.Text

		absoluteURL := e.Request.URL.String()
		descriptiveName := strings.Split(absoluteURL, "/3/")[1]
		filename := strings.Replace(descriptiveName, "/", "-", -1)
		filename = strings.Split(filename, ".")[0] + ".txt"
		writeDocToFile(filename, documentBody)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		absoluteLink := e.Request.AbsoluteURL(link)

		if e.Text == "next" {
			c.Visit(absoluteLink)
		}
	})

	c.Limit(&colly.LimitRule{
		 DomainGlob:  "*",
		 RandomDelay: 2 * time.Second,
 })

	// Start scraping on seed
  seed := "https://docs.python.org/3/tutorial/appetite.html"
	c.Visit(seed)
}
