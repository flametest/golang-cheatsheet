package main

import (
	"fmt"
	"time"
)

var LocationKualaLumpur, _ = time.LoadLocation("Asia/Kuala_Lumpur")

func main() {
	now := time.Now()
	future := time.Date(now.Year() + 1, now.Month(), now.Day(), 23, 59, 59, 0, LocationKualaLumpur)
	fmt.Println(future.Format(time.RFC3339))
}
