package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	nano := now.UnixNano()

	fmt.Println("Today : ", now)
	fmt.Println("Today's unix nano is : ", nano)

	UTCfromUnixNano := time.Unix(0, nano)

	fmt.Println("Today from Unix Nano : ", UTCfromUnixNano)
}