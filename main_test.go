package main


import "testing";
import "fmt";

	//webURL := "https://netflixtechblog.com/"
	//webURL := "https://tech.trello.com/"
	//webURL := "https://engineering.fb.com/"
	//webURL := "https://medium.com/pinterest-engineering"
	//webURL := "https://blog.cloudflare.com/"

func TestWebScraper(t *testing.T) {
	tws := []struct {
		name string
		url string
		result string
	} {
		{"netflix blog","https://netflixtechblog.com/","https://netflixtechblog.com/feed"},
		{"trello blog","https://tech.trello.com/","https://tech.trello.com/feed.xml"},
		{"Facebook blog","https://engineering.fb.com/","https://engineering.fb.com/feed/"},
		{"Pinterest blog","https://medium.com/pinterest-engineering","https://medium.com/feed/pinterest-engineering"},
		{"Cloudfare blog","https://blog.cloudflare.com","Cannot be retreived"},
		{"Y Combinator blog","https://news.ycombinator.com/","https://news.ycombinator.com/rss"},
		
	}

	for _, tt := range tws {
		test_name := fmt.Sprintf("%s",tt.name)
		t.Run(test_name,func (t *testing.T){
			ans := FindRSSFeed(tt.url)
			if(ans != tt.result) {
				t.Errorf("got %s, want %s",ans,tt.result)
			}
		})
	}
}