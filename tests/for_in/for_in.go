package main

import "fmt"

func main() {
	strings := []string{"hello", "world"}
	for idx, el := range strings {
		fmt.Println(idx, el)
		fmt.Println(idx, strings[idx])
	}
	for idx := range strings {
		fmt.Println(idx)
	}
	for range strings {
		fmt.Println("hello")
	}

	for i := range [2]int{} {
		for i := range [2]int{} {
			fmt.Println("hello")
		}
	}
}
