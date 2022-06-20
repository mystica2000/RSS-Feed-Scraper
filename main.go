package main

import (
	"fmt"
	"regexp"
	"github.com/gocolly/colly"
)

func matchForRSS(link string) bool {
	if(len(link)>0) {
	matched, _ := regexp.Match(`(feed|rss|feed.xml|Feed|RSS)`, []byte(link))
	comments, _ := regexp.Match(`(comments/feed)`, []byte(link))
	if(comments == true) {
		return false;
	}
	return matched;
	}
	return false
}

func FindRSSFeed(webURL string) string {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/6.0"),
		colly.IgnoreRobotsTxt(),
	)

	// if p, err := proxy.RoundRobinProxySwitcher(
	// ); err == nil {
	// 	c.SetProxyFunc(p)
	// }

	str := "Cannot be retreived";
	
	c.OnRequest(func(r *colly.Request){{
		fmt.Println("visiting", r.URL.String())
	}});

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "\nError:", err)
	})
	
	c.OnHTML("link[href]",func(h *colly.HTMLElement) {
		link := h.Attr("href")
		res := matchForRSS(link)
		if(res == true) {
		   fmt.Println("RSS FEED LINK = ",h.Request.AbsoluteURL(link))  
		   str = h.Request.AbsoluteURL(link)
		}
	});

	c.Visit(webURL)
	return str;
}


func main() {
	

	//webURL := "https://news.ycombinator.com/"
	webURL := "https://stackoverflow.blog"
	FindRSSFeed(webURL)
}
