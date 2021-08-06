package main

import (
	"encoding/json"
	"fmt"
)

type param struct {
	IDs []uint64
	Mm map[string]float64
}


func main() {
	p := &param{IDs: make([]uint64,0), Mm: map[string]float64{"xx": 234.1}}
	fmt.Println(p.IDs==nil)

	marshal, err := json.Marshal(p)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))
	s, err := json.Marshal(nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}
