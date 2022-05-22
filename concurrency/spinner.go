package concurrency

import (
	"fmt"
	"time"
)

func Spinner() {
	go spinner(100 * time.Millisecond)
	time.Sleep(time.Minute)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		fmt.Println("spinning")
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func terminator() {
	for {
		fmt.Println("Terminator terminated")
	}
}

func testing() {
	for {
		fmt.Println("Testing")
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
