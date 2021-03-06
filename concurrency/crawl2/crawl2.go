package crawl2

import (
	"fmt"
	"log"
	"mypkg/links"
	"os"
)

func Main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()
	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}
	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		log.Print(err)
	}
	return list
}

func Try() {
	numbers := make(chan int)
	for i := 0; i < 20; i++ {
		go func(i int) {
			numbers <- i
		}(i)
	}

	for num := range numbers {
		fmt.Println("num", num)
	}

	fmt.Println("done")
}
