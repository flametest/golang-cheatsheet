package main

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strings"
	"time"
)

var LocationKualaLumpur, _ = time.LoadLocation("Asia/Kuala_Lumpur")

type ActivationExtendInfo struct {

	// imei
	Imei string `json:"imei,omitempty"`

	// ktp name
	KtpName string `json:"ktpName,omitempty"`

	// ktp number
	KtpNumber string `json:"ktpNumber,omitempty"`

	// phone brand
	PhoneBrand string `json:"phoneBrand,omitempty"`

	// phone model
	PhoneModel string `json:"phoneModel,omitempty"`
}

func NumberInBetween(num, threshold1, threshold2 float64) bool {
	min := math.Min(threshold1, threshold2)
	max := math.Max(threshold1, threshold2)
	if num < min || num > max {
		return false
	} else {
		return true
	}
}

func main() {
	now := time.Now()
	future := time.Date(now.Year() + 1, now.Month(), now.Day(), 23, 59, 59, 0, LocationKualaLumpur)
	fmt.Println(future.Format(time.RFC3339))
	jsonDecoder := json.NewDecoder(strings.NewReader("{}"))
	extendInfo := &ActivationExtendInfo{}
	err := jsonDecoder.Decode(extendInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println(NumberInBetween(103.000000, 103, 119))
	s := decimal.NewFromFloat(43.00004).RoundUp(2).String()
	fmt.Println(s)
	var LocationManila, _ = time.LoadLocation("Asia/Manila")
	fmt.Println(LocationManila)
}
