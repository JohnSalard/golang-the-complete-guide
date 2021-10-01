package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://youtube.com",
		"http://amazon.com",
	}
	ch := make(chan string)
	for _, link := range links {
		go checkLink(link, ch)
		// fmt.Println(<-ch)
	}
	// fmt.Println(<-ch)
	for l := range ch {
		// fmt.Println("link: ", l)
		// go checkLink(l, ch)
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// ch <- "might be down!"
		ch <- link
		return
	}
	fmt.Println(link, "is up!")
	// ch <- "is up!"
	ch <- link
}
