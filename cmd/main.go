package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered in scrape", err)
		}
	}()
	// Request the HTML page.
	res, err := http.Get("https://apps.apple.com/cn/app/id414478124")
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	//供应商、应用大小、应用类别、价格、App 内购买项目、开发人员网站、App 支持、隐私政策
	supplier := doc.Find("dl > div:nth-child(1) > dd").Text()
	size := doc.Find(" dl > div:nth-child(2) > dd").Text()
	appType := doc.Find("dl > div:nth-child(3) > dd").Text()
	price := doc.Find("dl > div:nth-child(9) > dd").Text()
	appBuy := doc.Find("dl > div:nth-child(10) > dd > ol > div").Text()
	appSite, _ := doc.Find("div.small-hide.medium-show > ul > li:nth-child(1) > a").Attr("href")
	appSupply, _ := doc.Find("div.small-hide.medium-show > ul > li:nth-child(2) > a").Attr("href")
	secretPolicy, _ := doc.Find("div.small-hide.medium-show > ul > li:nth-child(3) > a").Attr("href")
	fmt.Println("供应商:", supplier)
	fmt.Println("应用大小:", size)
	fmt.Println("应用类别:", appType)
	fmt.Println("价格:", price)
	fmt.Println("App 内购买项目:", appBuy)
	fmt.Println("开发人员网站:", appSite)
	fmt.Println("App 支持:", appSupply)
	fmt.Println("隐私政策:", secretPolicy)
	wg.Done()
	return nil
}

var wg sync.WaitGroup

func main() {
	if len(os.Args)<2 {
		log.Fatal("命令行需要后面加参数，参数是数字")
	}
	num,_:=strconv.Atoi(os.Args[1])
	for i := 0; i < num; i++ {
		wg.Add(1)
		go ExampleScrape()
	}
	wg.Wait()
}
