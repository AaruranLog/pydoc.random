package main

import (
	"fmt"
	"time"
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
		if e.Text == "next" {
			c.Visit(absoluteLink)
		}
	})

	// On every a element which has span attribute call callback
	c.OnHTML("div.section", func(e *colly.HTMLElement) {
		fmt.Println("something new found")
		fmt.Println(e.ChildText("p"))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Limit(&colly.LimitRule{
		 DomainGlob:  "*",
		 RandomDelay: 1 * time.Second,
 })

	// Start scraping on seed
  seed := "https://docs.python.org/3/tutorial/appetite.html"
	c.Visit(seed)
}
