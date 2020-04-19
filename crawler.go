package main

import (
	// "database/sql"
	"net/url"
	"fmt"
	"github.com/gocolly/colly"
	_ "github.com/lib/pq"
	"os"
	"strings"
	"math/rand"
	"time"
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

func getRand(min int, max int) int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(max - min) + min
}

func crawl(file *os.File) {
	c := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"), colly.MaxDepth(3), colly.Async(true))
	outlink := ""
	input := "https://en.wikipedia.org/wiki/Slovakia"

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	cnt := 0
	c.OnHTML("#bodyContent a[href]", func(e *colly.HTMLElement) {
		inlink := e.Attr("href")
		outlink = e.Request.URL.String()[24:]
		cnt++

	 /* Tento big brain hack znizi pocet moznych hran v grafe na max zopar tisic tym, 
		* ze budem robit request iba raz za 50-100 moznych requestov.
		* Navyse sa bude vystupny subor vzdy lisit od toho predchadzajuceho.
		*/
		if n := getRand(50, 100); cnt % n != 0 {
			return
		}

		if strings.HasPrefix(inlink, "/wiki/") && 
		!strings.HasPrefix(inlink, "/wiki/Wikipedia") &&
		!strings.HasPrefix(inlink, "/wiki/Category") &&
		!strings.HasPrefix(inlink, "/wiki/File") &&
		!strings.HasPrefix(inlink, "/wiki/Help") &&
		!strings.HasPrefix(inlink, "/wiki/User") &&
		!strings.HasPrefix(inlink, "/wiki/Template") &&
		!strings.HasPrefix(inlink, "/wiki/ISBN") &&
		!strings.HasPrefix(inlink, "/wiki/Portal") && 
		!strings.HasPrefix(inlink, "/wiki/Special") {
			decodedOutlink, _ := url.QueryUnescape(outlink)
			decodedInlink, _ := url.QueryUnescape(inlink)
			fmt.Println(cnt, outlink, inlink)
			file.WriteString(decodedOutlink[6:] + "\t" + decodedInlink[6:] + "\n")
			e.Request.Visit(inlink)
		}
	})

	c.Visit(input)

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
