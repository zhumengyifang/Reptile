package Reptile

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

const blogUrl = "https://blog.csdn.net/weixin_40165163/article/list/"
const existUrl = "weixin_40165163"

func GetBlogVisitCount(url string) {
	if url == "" {
		url = blogUrl
	}

	ch := make(chan int, 10)
	request := make(chan int32)

	for i := 1; i <= 5; i++ {
		url := blogUrl + strconv.Itoa(i) + "?"
		fmt.Println("开始爬取：", url)
		go getPageBlogCount(url, ch)
	}

	go getResult(ch, request)
	fmt.Println("统计结果： ", <-request)
}

func getPageBlogCount(url string, ch chan int) {
	doc := getNewDoc(url)

	m, isLastUrl := getBlog(doc)
	if isLastUrl {

	}

	for k, v := range m {
		go getChildCount(k, v, ch)
	}
}

func getChildCount(url string, title string, ch chan int) {
	doc1 := getNewDoc(url)
	s, n := getVisitCount(doc1)
	fmt.Println("title:", title, "  ", s, "int:", n)
	ch <- n
}

func getNewDoc(url string) *goquery.Document {
	if url == "" {
		panic("Url Is Nil")
	}

	resp, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	if resp.StatusCode != 200 {
		panic("StatusCode:" + strconv.Itoa(resp.StatusCode))
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	return doc
}

func getResult(ch chan int, request chan int32) {
	time.Sleep(time.Second)
	var num int32 = 0
	for {
		select {
		case t := <-ch:
			atomic.AddInt32(&num, int32(t))
		default:
			request <- atomic.LoadInt32(&num)
		}
	}
}

func getBlog(doc *goquery.Document) (map[string]string, bool) {
	isLastUrl := true
	if doc == nil {
		panic("Doc Is Nil")
	}
	m := make(map[string]string)
	doc.Find(".article-list").Children().Each(func(i int, selection *goquery.Selection) {
		isLastUrl = false
		a := selection.Find("a").First()
		url, _ := a.Attr("href")
		title := a.Text()
		if strings.Contains(url, existUrl) {
			m[url] = title
		}
	})
	return m, isLastUrl
}

func getVisitCount(doc *goquery.Document) (string, int) {
	visitCount := doc.Find(".read-count").Text()
	return visitCount, getCount(visitCount, "：")
}

func getCount(visitCount string, subStr string) int {
	s := strings.Index(visitCount, subStr)
	n, err := strconv.Atoi(visitCount[s+len(subStr):])
	if err != nil {
		return 0
	}
	return n
}
