package main

import (
	// "database/sql"
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
	"os"
	"strings"
	"sync"
)

// const (
// 	host     = "ec2-34-204-22-76.compute-1.amazonaws.com"
// 	port     = 5432
// 	user     = "tttkiyickiurie"
// 	password = "e04c6810dad874a84873221afe0b1a7d95c7343558e2a4bc02297dd3e30aa374"
// 	dbname   = "d8q9700l99v1v8"
// )

// func connect() {
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=require",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// }

type safeURL struct {
	url string
	mux sync.Mutex
}

func crawl(file *os.File) {
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.MaxDepth(2), colly.Async(true))
	outlink := ""
	url := "https://en.wikipedia.org/wiki/Slovakia"

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		inlink := e.Attr("href")
		outlink = e.Request.URL.String()[24:]

		if strings.HasPrefix(inlink, "/wiki/") && strings.HasPrefix(outlink, "/wiki/") && inlink != outlink {
			fmt.Println(outlink, inlink)
			file.WriteString(outlink + "\t" + inlink + "\n")
		}

		e.Request.Visit(inlink)
	})

	c.Visit(url)

	c.Wait()
}

func main() {
	file, err := os.Create("links.txt")
	if err != nil {
		file.Close()
		return
	}
	crawl(file)
}
