package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func checkLink(link string, c chan<- string) {
	res, err := http.Get(link)
	if err != nil {
		log.Fatalf("failed to check %s: %v", link, err)
	}

	if res.StatusCode == 200 {
		c <- link + " is up"
		return
	}
	c <- link + " is down"
}

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.com",
		"https://ikon.mn",
		"https://news.gogo.mn",
	}
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(links))

	for _, link := range links {
		l := link
		go func() {
			checkLink(l, c)
			// fmt.Printf("%s status:%v\n", l, status)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for val := range c {
		fmt.Println(val)
	}

}
