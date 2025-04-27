package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

// Function to fetch the HTML content of a URL
func fetchHTML(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received non-200 status code %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	return doc, nil
}

// Function to extract the title from an HTML node
func extractTitle(doc *html.Node) string {
	var title string

	// Traverse the HTML tree to find the title tag
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			// Find the text content inside the <title> tag
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					title = c.Data
				}
			}
		}

		// Traverse through child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	return title
}

// Function to scrape the title from a list of URLs
func scrapeTitles(urls []string) {
	for _, url := range urls {
		// Fetch the HTML document from the URL
		doc, err := fetchHTML(url)
		if err != nil {
			log.Printf("Error fetching %s: %v\n", url, err)
			continue
		}

		// Extract the title from the document
		title := extractTitle(doc)

		// Display the result
		if title != "" {
			fmt.Printf("URL: %s\nTitle: %s\n\n", url, title)
		} else {
			fmt.Printf("URL: %s\nTitle: (No title found)\n\n", url)
		}
	}
}

func main() {
	// List of URLs to scrape
	urls := []string{
		"https://www.example.com",
		"https://golang.org",
		"https://www.github.com",
	}

	// Scrape titles from the provided URLs
	scrapeTitles(urls)
}
