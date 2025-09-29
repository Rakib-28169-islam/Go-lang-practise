package main

import(
	"fmt"
	"simple-web-scrape/Service"
)
func main(){
     urls:=[]string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.medium.com",
		"https://www.linkedin.com",
		"https://www.we.com",
	 }
	 results := Service.RequestParallel(urls)
	 for _, result := range results {
		fmt.Println(result)
	 }
}