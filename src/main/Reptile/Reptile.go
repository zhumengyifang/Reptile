package Reptile

import "C"
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

const blogUrl = "https://blog.csdn.net/weixin_40165163/article/list/"
const existUrl = "weixin_40165163"

func GetBlogVisitCount(url string) {
	if url == "" {
		url = blogUrl
	}

	isLastUrl := false
	for i := 1; !isLastUrl; i++ {
		url := blogUrl + strconv.Itoa(i) + "?"
		fmt.Println("开始爬取：", url)
		isLastUrl = GetPageBlogCount(url)
	}
}

func GetPageBlogCount(url string) bool {
	doc := GetNewDoc(url)
	m, isLastUrl := GetBlog(doc)
	if isLastUrl {
		return true
	}

	for k, v := range m {
		s, n := GetVisitCount(GetNewDoc(k))
		fmt.Println("title:", v, "  ", s, "int:", n)
	}
	return false
}

func GetNewDoc(url string) *goquery.Document {
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

func GetBlog(doc *goquery.Document) (map[string]string, bool) {
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

func GetVisitCount(doc *goquery.Document) (string, int) {
	visitCount := doc.Find(".read-count").Text()
	return visitCount, GetCount(visitCount, "：")
}

func GetCount(visitCount string, subStr string) int {
	s := strings.Index(visitCount, subStr)
	n, err := strconv.Atoi(visitCount[s+len(subStr):])
	if err != nil {
		return 0
	}
	return n
}
