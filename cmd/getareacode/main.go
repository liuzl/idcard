package main

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"github.com/liuzl/dl"
	"github.com/liuzl/goutil"
	"log"
	"strings"
	"unicode"
)

var (
	url = "http://www.mca.gov.cn/article/sj/xzqh/2018/201804-12/201804121005.html"
	//url       = "http://www.mca.gov.cn/article/sj/tjbz/a/201713/201708220902.html"
	rowsXpath = xpath.MustCompile(`//td[text()="行政区划代码"]/parent::tr/following-sibling::tr`)
)

func main() {
	resp := dl.DownloadUrl(url)
	if resp.Error != nil {
		log.Fatal(resp.Error)
	}
	doc, err := htmlquery.Parse(strings.NewReader(resp.Text))
	if err != nil {
		log.Fatal(err)
	}
	node := htmlquery.CreateXPathNavigator(doc)
	row, ok := rowsXpath.Evaluate(node).(*xpath.NodeIterator)
	if !ok {
		log.Fatal("rowsXpath.Evaluate(node).(*xpath.NodeIterator) error")
	}
	for row.MoveNext() {
		items := strings.Fields(row.Current().Value())
		if len(items) == 2 && goutil.StringIs(items[0], unicode.IsDigit) {
			fmt.Println(items)
		}
	}
}
