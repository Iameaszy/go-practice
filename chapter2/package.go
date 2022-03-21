package main

import (
	"fmt"

	"../person"
)

func main() {
	p := person.Person{
		Name: "Adeniyi Yusuf",
		Age: 26,
	}

	fmt.Println(p.Details())
}