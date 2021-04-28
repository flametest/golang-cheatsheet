package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	d1 := decimal.NewFromInt(19)
	d2 := decimal.NewFromInt(20)
	d := d1.Sub(d2)
	fmt.Println(d)
	// d1 does not change
	fmt.Println(d1)
}
