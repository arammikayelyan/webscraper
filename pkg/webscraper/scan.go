package webscraper

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// WebsiteInfo contains information about website
type WebsiteInfo struct {
	Website    string
	Links      []string
	StatusCode int
}

// getWebsiteInfo takes an array of websites and returns an information struct about them.
func GetWebsiteInfo(websites []string) []WebsiteInfo {
	var websiteInfos []WebsiteInfo
	// for each website call HTTP GET request
	for _, website := range websites {
		res, err := http.Get(website)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		defer res.Body.Close()

		// Check the status code of the response
		if res.StatusCode == http.StatusOK {
			doc, err := html.Parse(res.Body)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			var links []string
			var f func(*html.Node)

			// extract the link information and store it in the links array
			f = func(n *html.Node) {
				if n.Type == html.ElementNode && n.Data == "a" {
					for _, a := range n.Attr {
						if a.Key == "href" {
							links = append(links, a.Val)
							break
						}
					}
				}

				for c := n.FirstChild; c != nil; c = c.NextSibling {
					f(c)
				}
			}
			f(doc)

			websiteInfos = append(websiteInfos, WebsiteInfo{
				Website:    website,
				Links:      links,
				StatusCode: res.StatusCode,
			})
		} else {
			websiteInfos = append(websiteInfos, WebsiteInfo{
				Website:    website,
				StatusCode: res.StatusCode,
			})
		}
	}
	return websiteInfos
}
