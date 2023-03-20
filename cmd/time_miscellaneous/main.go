package main

import (
	"crypto/sha256"
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


type AeonOrder struct {
	OrderId         string    `json:"order_id"`
	ContactId       string    `gorm:"column:contact_id" json:"contact_id"`
	CustName        string    `gorm:"column:cust_id" json:"cust_name"`
	Gender          string    `gorm:"column:gender" json:"gender"`
	MobileTelp      string    `gorm:"column:mobile_telp" json:"mobile_telp"`
	Email           string    `gorm:"column:email" json:"email"`
	AgreementDate   time.Time `gorm:"column:agreement_date" json:"agreement_date"`
	EndDate         time.Time `gorm:"column:end_date" json:"end_date"`
	ProductPrice    int64     `gorm:"column:product_price" json:"product_price"`
	InsureAmount    int64     `gorm:"column:insure_amount" json:"insure_amount"`
	Category        string    `gorm:"column:category" json:"category"`
	Brand           string    `gorm:"column:brand" json:"brand"`
	NameAndType     string    `gorm:"column:name_and_type" json:"name_and_type"`
	ModelNo         string    `gorm:"column:model_no" json:"model_no"`
	Periode         string    `gorm:"column:periode" json:"periode"`
	IMEI            string    `gorm:"column:imei" json:"imei"`
	ReferenceNumber string    `json:"reference_number"`
	ConfirmChecked  string    `json:"confirm_checked"`
	BatchNumber     string    `json:"batch_number"`
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
	location, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		return
	}
	parse, err := time.ParseInLocation("1/2/2006", "4/1/2022", location)
	if err != nil {
		return
	}
	//2022-09-30T00:00:00+07:00
	fmt.Printf("%v", parse)

	orderStr := "{\"imei\": \"350331806261983\", \"brand\": \"Samsung\", \"email\": \"lilihsolihat323@gmail.com\", \"gender\": \"FEMALE\", \"periode\": \"\", \"category\": \"CO\", \"end_date\": \"2023-04-01T00:00:00+07:00\", \"model_no\": \"A53\", \"order_id\": \"6-02199739-6-1\", \"cust_name\": \"LILIH SOLIHAT\", \"contact_id\": \"6-02199739-6\", \"mobile_telp\": \"85772609588\", \"batch_number\": \"28042022\", \"insure_amount\": 290160, \"name_and_type\": \"HP SAMSUNG\", \"product_price\": 6200000, \"agreement_date\": \"2022-04-01T00:00:00+07:00\", \"confirm_checked\": \"\", \"reference_number\": \"\"}"
	order := &AeonOrder{}
	err = json.Unmarshal([]byte(orderStr), order)
	if err != nil {
		return
	}
	sourceStr := fmt.Sprintf("contractId:%s;productPrice:%v;brand:%s;model:%s;name:%s;startDate:%v;imei:%s", order.ContactId, order.ProductPrice, order.Brand, order.ModelNo, order.NameAndType, order.AgreementDate, order.IMEI)
	fmt.Println(sourceStr)
	hash := sha256.Sum256([]byte(sourceStr))
	externalId := fmt.Sprintf("%x", hash[:])
	fmt.Println(externalId)
}
