package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("docs.python.org", "http://docs.python.org",
		"https://docs.python.org", "https://docs.python.org/3/"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		absoluteLink := e.Request.AbsoluteURL(link)
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, absoluteLink)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// fmt.Printf("Body: %s\n", e.ChildText("body"))
		c.Visit(absoluteLink)
	})

	// On every a element which has span attribute call callback
	c.OnHTML("#section", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("p"))
		// class := e.Attr("class")
		// // Print link
		// fmt.Printf("Class: %s\n", class)
		// // Visit link found on page
		// // Only those links are visited which are in AllowedDomains
		// // fmt.Printf("Body: %s\n", e.ChildText("body"))
		// // c.Visit(absoluteLink)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on seed
  // seed := "https://godoc.org/"
  seed := "https://docs.python.org/3/tutorial/index.html"
	c.Visit(seed)
}
