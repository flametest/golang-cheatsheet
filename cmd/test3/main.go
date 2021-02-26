package main

import "fmt"

type Person struct {
	Addr *Address `json:"name"`
}
type Address struct {
	Street *Street  `json:"street"`
}
type Street struct {

}

func main() {
	p := &Person{Addr: &Address{}}
	fmt.Println(p.Addr.Street)
	if nil == p.Addr {
		fmt.Println("hello")
	}
 }
