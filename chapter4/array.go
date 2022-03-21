package main

import (
	"crypto/sha256"
	"fmt"
)


func main() {
	var a [3]int;
	var a2 [3]int = [3]int{1,2 ,3}
	a3  := [3]int{4,5, 6}
	a4  := [...]int{7}
	fmt.Printf("a2: %v\n",a2);
	fmt.Printf("a3: %v\n",a3);
	fmt.Printf("a4: %v\n",a4);
	fmt.Println(a[0]);
	fmt.Println(a);
	fmt.Println(len(a));
	fmt.Println(a[len(a) -1]);

	// for i := range a {
	// 	fmt.Println(i)
	// }

	c1 := sha256.Sum256([]byte("x")) 
	c2 := sha256.Sum256([]byte("X")) 
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}