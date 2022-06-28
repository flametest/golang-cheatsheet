package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	file := "/Users/jiangjun/go/src/github.com/axinan-com/goserver/axapp/elrond/email_templates/id-ID/buka_claim_expired_14days.html"
	ss := strings.Split(file, ".")
	fmt.Println(ss)
	s := ss[len(ss)-2]
	fmt.Println(s)
	ss = strings.Split(s, "/")
	fmt.Println(ss)
	notificationId := ss[len(ss)-1]
	fmt.Println(notificationId)
	fmt.Println(time.Now().UnixNano())
}
