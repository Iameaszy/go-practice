package greeting

import (
	"fmt"
	"time"
)

func Main() {
	done := make(chan bool)
	go func() {
		sayHello()
		done <- true
	}()
	fmt.Println("Hi")
	<-done
}

func sayHello() {
	time.Sleep(5 * time.Second)
	fmt.Println("Hello")
}
