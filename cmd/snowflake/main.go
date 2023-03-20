package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/sony/sonyflake"
	"strconv"
	"strings"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{})

func main() {
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	h := md5.New()
	h.Write([]byte(strconv.FormatUint(id, 10)))
	s := h.Sum(nil)
	hexValue := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz012345678912").EncodeToString(s)
	hexValue = strings.ToUpper(hexValue)
	fmt.Println(hexValue)
	fmt.Println(hexValue[:6])
}
