package main

import (
	"fmt"
	"time"
)

func main() {
	location, _ := time.LoadLocation("Asia/Shanghai")
	d1 := time.Date(2021, 1, 31, 9,2,3,12, location)
	d2 := time.Now()
	diff := d2.Sub(d1)
	fmt.Println(diff.Hours())
	fmt.Println(time.Now().Unix())
	millisecond := time.Now().UnixNano()/  int64(time.Millisecond)
	fmt.Println(millisecond)
	fmt.Println(millisecond / 1000)
	fmt.Println(1*1.0/2)
}
