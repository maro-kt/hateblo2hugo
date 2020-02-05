package transformer

import (
	"fmt"
  "regexp"
	"log"
  "net/url"
	"github.com/PuerkitoBio/goquery"
)

type MathTransformer struct {
	doc *goquery.Document
}

func (t *MathTransformer) Transform() error {

	t.doc.Find("img").Each(func(_ int, s *goquery.Selection) {
    src, _ := s.Attr("src")
    log.Printf("src = %s", src)
    rep := regexp.MustCompile(`https://chart.apis.google.com/chart\?cht=tx\&chl=(.*)`)
    str := rep.ReplaceAllString(src, "$1")
    log.Printf("regex: %s", str)
    str, _ = url.QueryUnescape(str)
    log.Printf("unescape: %s", str)
		s.ReplaceWithHtml(fmt.Sprintf("$%s$", str))
	})
	return nil
}
