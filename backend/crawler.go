package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"os"
	"strings"
)

func crawl(graph *graph) {
	c := colly.NewCollector(colly.AllowedDomains("awoiaf.westeros.org"), colly.MaxDepth(0), colly.Async(true))
	linkfile, _ := os.Create("linkfile.txt")
	outlink := ""
	cnt := 0
	input := "https://awoiaf.westeros.org/index.php/House_Targaryen"

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnHTML("#bodyContent a[href]", func(e *colly.HTMLElement) {
		inlink := e.Attr("href")
		outlink = e.Request.URL.String()
		cnt++

		/* Tento big brain hack znizi pocet moznych hran v grafe na max zopar tisic tym,
		* ze budem robit request iba raz za 50-100 moznych requestov.
		* Navyse sa bude vystupny subor vzdy lisit od toho predchadzajuceho.
		 */
		// if n := getRand(80, 81); cnt%n != 0 {
		// 	return
		// }

		/*
		* Pre istotu takto trapne exhaustivne, aby som nepridaval zbytocne stranky
		 */
		if strings.HasPrefix(inlink, "/index.php") &&
			strings.Contains(outlink, "Targaryen") &&
			strings.Contains(inlink, "Targaryen") &&
			!strings.HasPrefix(inlink, "/index.php/Wikipedia") &&
			!strings.HasPrefix(inlink, "/index.php?") &&
			!strings.Contains(inlink, "Chapter") &&
			!strings.Contains(inlink, "actor") &&
			!strings.Contains(inlink, "Calculations") &&
			!strings.Contains(inlink, "TV") &&
			!strings.HasPrefix(inlink, "/index.php/Category") &&
			!strings.HasPrefix(inlink, "/index.php/File") &&
			!strings.HasPrefix(inlink, "/index.php/Help") &&
			!strings.HasPrefix(inlink, "/index.php/User") &&
			!strings.HasPrefix(inlink, "/index.php/Template") &&
			!strings.HasPrefix(inlink, "/index.php/ISBN") &&
			!strings.HasPrefix(inlink, "/index.php/Portal") &&
			!strings.HasPrefix(inlink, "/index.php/Talk") &&
			!strings.HasPrefix(inlink, "/index.php/Special") {
			decodedOutlink, _ := url.PathUnescape(outlink)
			decodedInlink, _ := url.PathUnescape("https://awoiaf.westeros.org" + inlink)
			fmt.Println(cnt, decodedOutlink, decodedInlink)
			linkfile.WriteString(decodedOutlink + "\t" + decodedInlink + "\n")
			graph.link(decodedOutlink, decodedInlink, 1.0)
			e.Request.Visit(inlink)
		}

	})

	c.Visit(input)

	c.Wait()
}
