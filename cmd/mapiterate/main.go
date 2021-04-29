package main

import "fmt"

func main() {
	m := map[string]uint64{
		"x": 1,
	}
	for s, u := range m {
		fmt.Println(s, u)
	}
}
