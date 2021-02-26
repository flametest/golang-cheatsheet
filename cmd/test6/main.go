package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func init()  {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	res := make([]rune, 6)
	for i := 0; i < 6; i++ {
		idx := rand.Intn(len(letters))
		res[i] = rune(letters[idx])
	}
	fmt.Println(string(res))
}
