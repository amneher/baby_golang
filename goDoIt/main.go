package main

import (
	"fmt"
)

func main() {
	var s = make([]string, 5)
	values := []string{"one", "two", "three", "four", "five"}
	for i := range s {
		s[i] = values[i]
	}
	fmt.Println(s)
}
