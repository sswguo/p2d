package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-git/go-git/v5"
	"github.com/robfig/cron"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Repo struct {
	Name  string
	Link  string
	Desc  string
	Lan   string
	Stars string
	Forks string
}

func trending() {

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	
	// request html content
	resp, err := client.Get("https://github.com/trending")

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()

	doc, err1 := goquery.NewDocumentFromReader(resp.Body)
	if err1 != nil {
		log.Fatal(err1)
	}

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		if i > 2 {
			return
		}
		repo := Repo{}
		fmt.Printf("--------")
		fmt.Printf("\n")

		// description
		desc := strings.TrimSpace(s.Find("p.col-9.color-text-secondary.my-1.pr-4").Text())
		//fmt.Printf(desc)
		repo.Desc = desc

		// link
		link, _ := s.Find(".lh-condensed a").Attr("href")
		linkStr := strings.TrimSpace(link)
		repo.Link = "https://github.com" + linkStr
		repo.Name = linkStr[1:]

		// main language
		lan := s.Find("span.d-inline-block.ml-0.mr-3 > span:nth-child(2)").Text()
		repo.Lan = lan

		s.Find("a.Link--muted.d-inline-block.mr-3").Each(func(j int, s2 *goquery.Selection) {
			num := strings.TrimSpace(s2.Text())
			if j == 0 {
				repo.Stars = num
			} else if j == 1 {
				repo.Forks = num
			}
		})
		fmt.Printf("Repo info: %+v\n", repo)
		fmt.Printf("\n")
	})

	_, err2 := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/freeCodeCamp/freeCodeCamp",
		Progress: os.Stdout,
	})

	if err2 != nil {
		fmt.Print(err2)
	}

	_, err4 := os.Stat("/tmp/foo/README.md")
	if os.IsNotExist(err4) {
		return
	}

	b, err3 := ioutil.ReadFile("/tmp/foo/README.md")
	if err3 != nil {
		fmt.Print(err3)
	}

	str := string(b)
	fmt.Println(str)

}

func main() {

	i := 0
	c := cron.New()
	spec := "0 */1 * * * *"
	c.AddFunc(spec, func() {
		i++
		fmt.Println("execute per minute", i)
		trending()
	})
	c.Start()
	select {}
}
