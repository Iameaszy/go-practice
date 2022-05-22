package main

import (
	"fmt"
	"sort"
)


func init() {
	var person map[string]string = map[string]string{
		"firstName": "Yusuf",
		"lastName": "Adeniyi",
		"age": "21",
		"height": "5 meters",
	}
	fmt.Println("person: ", person["age"]);

	person2 := make(map[string]string)
	person2["firstName"] = "Yusuf";
	person2["lastName"] = "Adeniyi";
	delete(person2, "firstName");
	fmt.Println("person2: ", person2);

	for key, value := range person {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	keys := make([]string, 0, len(person))
	for key := range person {
		keys = append(keys, key)
	}
	sort.Strings(keys);
	fmt.Println("keys", keys);

	if value, ok := person["height"]; ok {
		fmt.Println("height", value)
	}

	// True if equal is written incorrectly. 
	isEqual := equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Println("isEqual", isEqual);

	fmt.Println("complexion", person["complexion"])
}

func equal(x, y map[string]int) bool { 
	if len(x) != len(y) { 
		return false 
	} 
	for k, xv := range x { 
		if yv, ok := y[k]; !ok || yv != xv { 
			fmt.Printf("k: %v, xv: %v, yv: %v", k, xv, yv)
			return false 
		} 
	} 
	return true 
}