package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

func main() {
	fmt.Println(decimal.NewFromFloat(float64(10) / 100))
	var d decimal.Decimal
	fmt.Println(d.String())
	planIDSet := make(map[uint64]bool, 0)
	planIDSet[1] = true
	fmt.Println(len(planIDSet))
	fmt.Println("xx", planIDSet)
	for u, b := range planIDSet {
		fmt.Println(u, b)
	}
	var latestOrderTime time.Time
	fmt.Println(latestOrderTime)

	var a int64
	fmt.Println(a)
}