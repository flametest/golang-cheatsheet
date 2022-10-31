package main

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkReadAll(b *testing.B) {
	file, err := os.Open("/Users/jiangjun/go/src/github.com/flametest/golang-cheatsheet/cmd/ioReadTest/test.pdf")
	if err != nil {
		fmt.Errorf("open err:%v", err)
		return
	}
	for n := 0; n < b.N; n++ {
		IOReadAll(file) // run fib(30) b.N times
	}
}

func BenchmarkCopy(b *testing.B) {
	file, err := os.Open("/Users/jiangjun/go/src/github.com/flametest/golang-cheatsheet/cmd/ioReadTest/test.pdf")
	if err != nil {
		fmt.Errorf("open err:%v", err)
		return
	}
	for n := 0; n < b.N; n++ {
		IOCopy(file) // run fib(30) b.N times
	}
}
