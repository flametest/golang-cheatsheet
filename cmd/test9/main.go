package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n float64
	n = 34.222
	s := strconv.FormatFloat(n, 'f', -1, 64)
	fmt.Println(s)
	s1 := strconv.FormatFloat(n, 'E', -1, 64)
	fmt.Println(s1)
	a := []int{1,2,3,4,5}
	b := a[0:1000]
	fmt.Println(b)
}
