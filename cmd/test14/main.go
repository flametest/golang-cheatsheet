package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS(":9090", "/Users/jiangjun/go/src/github.com/flametest/golang-cheatsheet/cmd/test14/server.crt",
		"/Users/jiangjun/go/src/github.com/flametest/golang-cheatsheet/cmd/test14/server.key", nil); err != nil {
		fmt.
			Println(err)
	}
}