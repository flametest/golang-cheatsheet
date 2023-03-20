package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println(decimal.NewFromFloat(1.104).Round(2).String())
	fmt.Println(decimal.NewFromFloat(1.105).Round(2).String())
}
