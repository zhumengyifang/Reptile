package Reptile

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

//目标网址
const blogUrl = "https://blog.csdn.net/weixin_40165163/article/list/"
const existUrl = "weixin_40165163"

func GetBlogVisitCount(url string) {
	if url == "" {
		url = blogUrl
	}

	//统计总的访问量
	count := 0
	//判断是否为最后一页
	isLastUrl := false
	for i := 1; !isLastUrl; i++ {
		n := 0
		//循环拼下一页的url
		url := blogUrl + strconv.Itoa(i) + "?"
		fmt.Println("开始爬取：", url)
		//按页去统计访问量count
		isLastUrl, n = GetPageBlogCount(url)
		//汇总
		count += n
	}
	fmt.Println("总访问量：", count)
}

func GetPageBlogCount(url string) (bool, int) {
	//获取需要爬取url的document对象
	doc := GetNewDoc(url)

	//获取当前页面所有博客的连接和title
	m, isLastUrl := GetBlog(doc)
	//如果当前页面爬不到url则为最后一页结束循环
	if isLastUrl {
		return true, 0
	}

	count := 0
	for k, v := range m {
		//获取详细博客内容中的访问量
		s, n := GetVisitCount(GetNewDoc(k))
		fmt.Println("title:", v, "  ", s, "int:", n)
		count += n
	}
	return false, count
}

//获取url地址的document对象
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

//获取doc对象中的url和title
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

//解析 访问量：123 中的 123
func GetCount(visitCount string, subStr string) int {
	s := strings.Index(visitCount, subStr)
	n, err := strconv.Atoi(visitCount[s+len(subStr):])
	if err != nil {
		return 0
	}
	return n
}
