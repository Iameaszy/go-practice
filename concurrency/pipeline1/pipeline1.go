package pipeline1

import "fmt"

func Main() {
	naturals := make(chan int)
	squares := make(chan int)
	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()
	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()
	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}