package main

import (
	"crypto/sha256"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Test struct {
	Debug bool
}

func PhoneSignature(androidID, uuid string) string {
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s:%s", androidID, uuid)))
	return fmt.Sprintf("%x", sum)
}

func main() {
	t := Test{}
	fmt.Println(t.Debug)
	fmt.Println(PhoneSignature("e1cdd94a89dcfb3b", ""))
	id := uuid.NewV4().String()
	fmt.Println(id, len(id))
	var LocationKualaLumpur, _ = time.LoadLocation("Asia/Kuala_Lumpur")
	t1 := time.Date(2021, 8, 18, 12,0,0, 0, LocationKualaLumpur)
	fmt.Println(t1.Add(time.Hour * 24 * 14).String())
}
