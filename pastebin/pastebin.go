package pastebin

import (
	"fmt"
	"paste-corral/data"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

const (
	baseURL string = "https://pastebin.com"
	apiURL  string = baseURL + "/api/"
)

// GetArchiveURLs gets the list of all available archive URLs.
func GetArchiveURLs() (urls []string) {

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("pastebin.com"),
	)

	// Append child archive href to urls slice
	c.OnHTML("div.lang_div", func(e *colly.HTMLElement) {

		ch := e.DOM.Children()
		link, _ := ch.Eq(0).Children().Eq(0).Attr("href")

		urls = append(urls, link)
	})

	// Start scraping:
	c.Visit(apiURL)

	return urls
}

// GetPasteURLs gets the list of all URLs for a single archive.
func GetPasteURLs(archiveURL string) (urls []string) {

	c := colly.NewCollector(
		colly.AllowedDomains("pastebin.com"),
		colly.AllowURLRevisit(),
	)

	// Add random delay of a maximum of 5 seconds
	c.Limit(&colly.LimitRule{
		RandomDelay: 5 * time.Second,
	})

	// Generate a new User-Agent string before every request
	extensions.RandomUserAgent(c)

	// -- --------------------------------------
	// Append child archive href to urls slice
	c.OnHTML("tr", func(e *colly.HTMLElement) {

		ch := e.DOM.Children()
		eClass, _ := ch.Eq(0).Children().Eq(0).Attr("class")

		// incorrect class
		if eClass != "i_p0" {
			return
		}

		link, _ := ch.Eq(0).Children().Eq(1).Attr("href")

		urls = append(urls, link)
	})

	// Start scraping:
	c.Visit(baseURL + archiveURL)

	return urls

}

// GetPaste gets the actual pastebin post.
func GetPaste(pasteURL string) (author, title, content, dt string) {

	time.Sleep(2 * time.Minute)

	c := colly.NewCollector(
		colly.AllowedDomains("pastebin.com"),
		colly.AllowURLRevisit(),
	)

	// Add random delay of a maximum of 7 seconds
	c.Limit(&colly.LimitRule{
		RandomDelay: 7 * time.Second,
	})

	// Generate a new User-Agent string before every request
	extensions.RandomUserAgent(c)

	// -- --------------------------------------
	// Get author:
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		attrSubstr := `/message_compose?to=`
		attr := e.Attr("href")

		// incorrect attr
		if !strings.HasPrefix(attr, attrSubstr) {
			return
		}

		author = strings.Replace(attr, attrSubstr, "", -1)

	})

	// -- --------------------------------------
	// Get title:
	c.OnHTML("div.paste_box_line1", func(e *colly.HTMLElement) {
		title = e.Text
	})

	// -- --------------------------------------
	// Get content:
	c.OnHTML("textarea.paste_code", func(e *colly.HTMLElement) {
		content = e.Text
	})

	// -- --------------------------------------
	// Get date:
	c.OnHTML("div.paste_box_line2", func(e *colly.HTMLElement) {

		pattern := `[Ss]unday|[Mm]onday|[tT]uesday|[Ww]ednesday|[Tt]hursday|[Ff]riday|[Ss]aturday`

		ch := e.DOM.Children()

		attr, _ := ch.Eq(4).Attr("title")

		_, err := regexp.MatchString(attr, pattern)

		// incorrect attr
		if err != nil {
			return
		}

		dt = attr

	})

	// Start scraping:
	c.Visit(baseURL + pasteURL)

	return author, title, content, dt

}

// Crawl starts the crawler which runs in infinite loop.
func Crawl() {

	// Infinite loop:
	for {

		archiveURLs := GetArchiveURLs()

		// -- -----------------------------------
		for _, archiveURL := range archiveURLs {

			pasteURLs := GetPasteURLs(archiveURL)

			// -- -------------------------------
			for _, pasteURL := range pasteURLs {

				author, title, content, dt := GetPaste(pasteURL)

				err := data.CreateRawPaste(author, title, content, dt)

				if err != nil {
					fmt.Println(err)
				}

			}

		}

		time.Sleep(2 * time.Minute)

	}
}
