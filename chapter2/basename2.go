package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("basename: %v", basename("a/b.c.go"))
	fmt.Printf("comma: %s", comma("1234"))
	s := "abc"
	b := []byte(s);
	s2 := string(b)

	fmt.Printf("s: %s, b: %v, s2: %s\n", s, b, s2)
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"

	for i,v := range "hello" {
		fmt.Printf("{index: %d, value: %c}\n", i, v);
	}

	x := 123 
	y := fmt.Sprintf("%d", x)
	xLength := len("hello");
	xLength = 3;
	fmt.Println("xLength: %d\n", xLength);
	fmt.Println(y, strconv.Itoa(x)) // "123 123"
	fmt.Println(strconv.FormatInt(int64(x), 9)) // "1111011"


	const (
		e = 1 + iota
		f 
		g 
		h
	)
	fmt.Println("e, f, g, h: %v, %v, %v, %v \n", e,f,g,h)
}

func basename(s string) string { 
	slash := strings.LastIndex(s, "/") // -1 if "/" not found 
	s = s[slash+1:] 
	if dot := strings.LastIndex(s, "."); dot >= 0 { 
		s = s[:dot] 
	} 
	return s 
}

func comma(s string) string { 
	n := len(s) 
	if n <= 3 { return s } 
	return comma(s[:n-3]) + "," + s[n-3:] 
}

func intsToString(values []int) string { 
	var buf bytes.Buffer 
	buf.WriteByte('[') 
	for i, v := range values { 
		if i > 0 { 
			buf.WriteString(", ") 
		} 
		fmt.Fprintf(&buf, "%d", v) 
	} 
	buf.WriteByte(']') 
	return buf.String() 
}