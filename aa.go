package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

func main() {
  req, _ := http.NewRequest("GET", "http://yahoo.co.jp", nil)
  req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36")
  res, _ := new(http.Client).Do(req)

  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
  }

  doc, err := goquery.NewDocumentFromReader(res.Body)
  if err != nil {
    log.Println(err)
  }

  li := doc.Find("main article section ul").Eq(0).Find("li")
  li.Each(func(index int, s *goquery.Selection) {
    fmt.Println(s.Text())

    tmp, err := s.Find("a").Attr("href")
    if err != true {
      log.Fatal(err)
    }
    fmt.Println(tmp + "\n")
  })
}
