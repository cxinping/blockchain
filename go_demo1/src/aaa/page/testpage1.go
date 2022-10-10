package page

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func TestHttp1() {
	html := `<body>
				<div>DIV1</div>
				<div>DIV2</div>
				<div>DIV3</div>
				<span>SPAN</span>
			</body>
			`
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})

}
