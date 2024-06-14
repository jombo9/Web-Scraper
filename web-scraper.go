package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/PuerkitoBio/goquery"
)

func main() {
    // Send an HTTP request to the target website.
    res, err := http.Get("https://example.com")
    if err != nil {
        log.Fatal(err)
    }
    defer res.Body.Close()

    // Check if the HTTP request was successful.
    if res.StatusCode != 200 {
        log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }

    // Load the HTML document.
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        log.Fatal(err)
    }

    // Find and print the page title.
    doc.Find("title").Each(func(i int, s *goquery.Selection) {
        title := s.Text()
        fmt.Println("Page Title: ", title)
    })
}
