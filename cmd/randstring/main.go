package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	rand.Seed(time.Now().Unix())
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	s, err := GenerateRandomString(16)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
}


