package http

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func TestHttp1() {
	response, _ := http.Get("http://www.baidu.com")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestHttp2() {
	response, err := http.Get("http://www.baidu.com/")

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", string(contents))
	}

}

func TestHttp3() {
	resp, err := http.Get("http://www.baidu.com")
	// resp, err := http.Get("http://www.163.com")
	if err != nil {
		fmt.Println("http get error.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error")
	}

	src := string(body)

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}

func TestHttp4() {
	res, err := http.Get("https://www.zol.com.cn/")
	if err != nil {
		// 错误处理
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)

	}
	fmt.Println("-------------------------------")
	fmt.Println(res.Body)
}

func TestHttp5() {
	var num = 0
	fmt.Println("请输入数字")
	fmt.Scanln(&num)
	fmt.Println("num=>", num)

	url := "https://www.zol.com.cn/"
	result, err := HttpGet(url)
	fmt.Println(result, err)
}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//循环读取网页数据
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加每一次循环读取到的数据
		result += string(buf[:n])
	}
	return result, err
}

func TestHttp6() {
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
		fmt.Println(i, selection.Text())
	})
}

func Example1() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}

func Example2() {
	c := colly.NewCollector()

	// selector goquery name id class
	c.OnHTML(".sidebar-link", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))

		ret, _ := e.DOM.Html()
		fmt.Println("ret-> ", ret)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("url => ", r.URL)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://gorm.io/zh_CN/docs")
}
