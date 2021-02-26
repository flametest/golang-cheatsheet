package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

type Person struct{

}

type EntityTime time.Time
func main() {
	var m = map[string]*Person{
		"1": &Person{},
		"2": &Person{},
	}
	fmt.Println(m["2"])
	fmt.Println(m["3"])
	uid := uuid.NewV4().String()
	fmt.Println(uid[len(uid) - 6:])
	fmt.Println(newCode(uid[len(uid) - 6:]))
}

func newCode(uuid string) string {
	chars := []rune(uuid)
	fmt.Println(chars)
	for index, value := range chars {
		if value > 'z' || value < 'a' {
			continue
		}
		if rand.Int()%2 == 0 {
			chars[index] -= 32
		}
	}
	return string(chars)
}
