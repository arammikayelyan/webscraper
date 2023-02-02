package main

import (
	"fmt"
	"github.com/arammikayelyan/webscraper/pkg/webscraper"
)

func main() {
	websites := []string{
		"https://www.webscraper.io/test-sites/e-commerce/ajax",
		"https://codedamn-classrooms.github.io/webscraper-python-codedamn-classroom-website/",
	}
	websiteInfos := webscraper.GetWebsiteInfo(websites)

	for i, winfo := range websiteInfos {
		fmt.Printf("Website %d: \n", i+1)
		fmt.Printf("\tWebsite name: %s \n", winfo.Website)
		fmt.Printf("\tStatus code: %d \n", winfo.StatusCode)
		fmt.Printf("\tLinks: \n")
		for j, link := range winfo.Links {
			fmt.Printf("\t\tLink %d: %s\n", j+1, link)
		}
		fmt.Println()
	}
}
