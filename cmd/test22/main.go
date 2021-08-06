package main

import "fmt"

type Test struct {
	Debug bool
}

func main() {
	t := Test{}
	fmt.Println(t.Debug)
}
