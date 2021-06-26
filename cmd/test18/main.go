package main

import (
	"fmt"
	"math/rand"
	"time"
)

const letters = "ABCDEFGHJKMNPQRSTUVWXYZabcdefghjkmnpqrstuvwxyz0123456789"
const A = "igloo_app_psp"
const B = "shopeemy"

func main() {
	fmt.Println(GenerateRandomString(8))
	appName := "igloo_app_psp"
	fmt.Println(SSS(appName))
}

func SSS(appName string) string {

	switch appName {
	case A:
	case B:

		return "xx"
	}
	return "yy"
}

func GenerateRandomString(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, 0)
	for i := 0; i < n; i++ {
		idx := rand.Intn(len(letters))
		b = append(b, letters[idx])
	}
	return string(b)
}