package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func main() {
	m := md5.New()
	m.Write([]byte("ไพรัช เพลินมาลัย"))
	fmt.Println(base64.StdEncoding.EncodeToString(m.Sum(nil)))
	fmt.Println()
}
