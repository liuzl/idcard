package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"crawler.club/dl"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"zliu.org/goutil"
)

var (
	rowsXpath = xpath.MustCompile(`//td[text()="行政区划代码"]/parent::tr/following-sibling::tr`)
	dataPath  = `../../data`
)

func CodeMap(url string) (map[string]string, error) {
	resp := dl.DownloadUrl(url)
	if resp.Error != nil {
		return nil, resp.Error
	}
	doc, err := htmlquery.Parse(strings.NewReader(resp.Text))
	if err != nil {
		return nil, err
	}
	node := htmlquery.CreateXPathNavigator(doc)
	row, ok := rowsXpath.Evaluate(node).(*xpath.NodeIterator)
	if !ok {
		return nil, fmt.Errorf("rowsXpath.Evaluate(node).(*xpath.NodeIterator) error")
	}
	m := make(map[string]string)
	for row.MoveNext() {
		items := strings.Fields(row.Current().Value())
		if len(items) == 2 && goutil.StringIs(items[0], unicode.IsDigit) {
			m[items[0]] = items[1]
		}
	}
	return m, nil
}

func main() {
	for _, s := range sources {
		log.Printf("crawl %s", s.Url)
		m, err := CodeMap(s.Url)
		if err != nil {
			log.Fatal(err)
		}
		b, err := goutil.JSONMarshalIndent(m, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		path := filepath.Join(dataPath, s.ID+".json")
		if err = ioutil.WriteFile(path, b, os.FileMode(0664)); err != nil {
			log.Fatal(err)
		}
	}
}
