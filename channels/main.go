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
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// infinity loop in go
	// for {
	// 	go checkLink(<-c, c)
	// }

	// alternative way of above is
	for l := range c {
		// function literal which same as anonymus function in javascript
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Is down")
		c <- link
		return
	}
	c <- link
	fmt.Println(link, "is up!")
}
