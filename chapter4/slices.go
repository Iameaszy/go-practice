package main

import (
	"fmt"
)


func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5:"May", 6:"June", 7:"July", 8:"August", 9: "September", 10: "October", 11: "November",12: "December"}
	fmt.Printf("months:%v, len:%d, cap: %d\n", months, len(months), cap(months));
	firstSixMonths := months[4:7]
	fmt.Printf("first six months: %v, len: %d, cap: %d, test: %v\n", firstSixMonths, len(firstSixMonths), cap(firstSixMonths), firstSixMonths[:9]);
	fmt.Printf("months[0]: %v\n", months[0]);

	
	mixedArr := [2][2]int{{1,2}, {}}
	fmt.Println("mixedArr", mixedArr)
	arr := []int{1, 2, 3, 4, 5, 6}
	reverse(arr);
	arr = append(arr, 7);
	fmt.Println("arr", arr)

	arr2 := make([]int, 3, 5)
	fmt.Printf("arr2: %v, len: %d, cap: %d\n", arr2, len(arr2), cap(arr2))


	var runes []rune
	for _,j := range "hello" {
		runes = append(runes, j)
	}
	fmt.Printf("runes %c, len: %d\n", runes, len(runes));
	fmt.Printf("hello in rune: %s\n", []byte("hello"))

	var x, y []int 
	for i := 0; i < 10; i++ { 
		y = appendInt(x, i) 
		fmt.Printf("%d cap=%d\t%d\t%v\n", i, cap(y), len(y), y) 
		x = y 
	}
	
	numbers := []int{1,4,3,3,4}
	numbers = append(numbers,3)
	fmt.Printf("cap=%d\t%d\t%v\n", cap(numbers), len(numbers), numbers); 
	fmt.Printf("numbers %v\n", append(numbers,x...)); 

	var word []string = []string{"one","","three"}
	fmt.Printf("len: %d, cap: %d, value: %s\n", len(word), cap(word), word);
	var wordWithoutEmpty []string = nonempty2(word);
	fmt.Printf("len: %d, cap: %d, value: %s\n", len(wordWithoutEmpty), cap(wordWithoutEmpty), wordWithoutEmpty);
}


func reverse(s []int) { 
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 { 
		s[i], s[j] = s[j], s[i] 
	} 
}

func appendInt(x []int, y int) []int { 
	var z []int 
	zlen := len(x) + 1 
	if zlen <= cap(x) { 
		// There is room to grow. Extend the slice. 
		z = x[:zlen] 
	} else { 
		// There is insufficient space. Allocate a new array. 
		// Grow by doubling, for amortized linear complexity. 
		zcap := zlen 
		if zcap < 2*len(x) { 
			zcap = 2 * len(x) 
		} 
		z = make([]int, zlen, zcap) 
		copy(z, x) // a built-in function; see text 
	}
 	z[len(x)] = y 
	return z 
}

// nonempty returns a slice holding only the non-empty strings. 
// The underlying array is modified during the call. 
func nonempty(strings []string) []string { 
	i := 0 
	for _, s := range strings { 
		if s != "" { 
			strings[i] = s 
			i++ 
		} 
	} 
	return strings[:i];
}

func nonempty2(strings []string) []string { 
	out := strings[:0] 
	// zero-length slice of original 
	for _, s := range strings { 
		if s != "" { 
			out = append(out, s) 
		} 
	} 
	return out 
}

func remove(slice []int, i int) []int { 
	copy(slice[i:], slice[i+1:]) 
	return slice[:len(slice)-1] 
}

func remove1(slice []int, i int) []int { 
	slice[i] = slice[len(slice)-1] 
	return slice[:len(slice)-1] 
}