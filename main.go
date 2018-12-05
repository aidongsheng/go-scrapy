package main

import (
	"github.com/aidongsheng/go-scrapy/sse"
	"log"
)

//import (
//	"fmt"
//	"github.com/PuerkitoBio/goquery"
//	"log"
//)
//
//func ExampleScrape() {
//	doc, err := goquery.NewDocument("http://www.qiushibaike.com")
//	if err != nil {
//		log.Fatal(err)
//	}
//	doc.Find(".article").Each(func(i int, s *goquery.Selection) {
//		if s.Find(".thumb").Nodes == nil && s.Find(".video_holder").Nodes == nil {
//			content := s.Find(".content").Text()
//			fmt.Printf("%s", content)
//
//		}
//	})
//}
//
//func main() {
//	ExampleScrape()
//}

func main() {

	var err error
	err = sse.GetStockListA("/Users/josan/Documents/sseA.csv")
	if err != nil {
		log.Fatal(err)
	}
	err = sse.GetStockListB("/Users/josan/Documents/sseB.csv")
	if err != nil {
		log.Fatal(err)
	}

}
