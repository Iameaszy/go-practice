package main

import (
	"fmt"
	"os"
)


 func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i];
		sep += " "
	}

	for index, arg := range os.Args[1:] {
		println("arg: ", arg, "index: ", index)
	}


	// fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args)
 }