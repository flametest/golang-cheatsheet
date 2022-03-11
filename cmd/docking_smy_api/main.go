package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	ShopeeMYReqAppId = "igloo"
	AppSecretKey = "f4ca48506a31a72c007e665244b0b3d3"
	ShopeeMYClaimHost = "https://insure-shopee.qa.iglooinsure.com"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	randStr := base64.URLEncoding.EncodeToString(b)
	fmt.Println(randStr)
	return randStr, err
}

func GenerateSignature(appSecretKey, method, url, timestampStr, nonceStr, bodyStr string) (string, error) {
	sigStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", method, url, timestampStr, nonceStr, bodyStr)
	fmt.Println(sigStr)
	h := hmac.New(sha256.New, []byte(appSecretKey))
	_, err := h.Write([]byte(sigStr))
	if err != nil {
		return "", err
	}
	sig := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(sig), nil
}

type SigTripper struct {
	Proxies http.RoundTripper
}

func (t SigTripper) RoundTrip(req *http.Request) ( res *http.Response,e  error) {
	req.Header.Set("X-Req-AppId", ShopeeMYReqAppId)
	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	req.Header.Set("X-Req-Timestamp", timestampStr)
	nonceStr, err := GenerateRandomString(16)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Req-Nonce", nonceStr)
	method := req.Method
	url := req.URL.Path
	buf, _ := ioutil.ReadAll(req.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

	req.Body = rdr1
	body, err := ioutil.ReadAll(rdr2)
	if err != nil {
		return nil, err
	}
	bodyStr := string(body)
	fmt.Println(bodyStr)
	sig, err := GenerateSignature(AppSecretKey, method, url, timestampStr, nonceStr, bodyStr)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Req-Signature", sig)
	res, e = t.Proxies.RoundTrip(req)
	if e != nil {
		log.Error(e)
	}
	return
}
type NotifyClaimInfoReq struct {
	ShopeePolicyNO string `json:"shopee_policy_no"`
	ClaimNO string `json:"claim_no"`
	PolicyNO string `json:"policy_no"`
	ActionTime string `json:"action_time"`
	ClaimAmount string `json:"claim_amount"`
	ClaimAmountApproved string `json:"claim_amount_approved"`
	Status string `json:"status"`
	Message string `json:"message"`
}

type NotifyClaimInfoV2 struct {
	ClaimNotificationNo string `json:"claim_notification_no"`
	LossOccurredType string `json:"loss_occurred_type"`
	LossOccurredTime string `json:"loss_occurred_time"`
	Policies []*RefPolicyInfo `json:"policies"`
	ActionTime string `json:"action_time"`
	Status string `json:"status"`
	Message string `json:"message"`
}

type RefPolicyInfo struct {
	ShopeePolicyNo string `json:"shopee_policy_no"`
	PolicyNo string `json:"policy_no"`
	Currency string `json:"currency"`
}

type UnderwritingParam struct {
	ShopeePolicyNo    string `json:"shopee_policy_no"`
	RefApplicationNo  string `json:"ref_application_no"`
	PolicyNo          string `json:"policy_no"`
	PolicyContractUrl string `json:"policy_contract_url"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	Status            string `json:"status"`
	Message           string `json:"message"`
	ClaimUrl          string `json:"claim_url"`
}

type PolicyLapseParam struct {
	ShopeePolicyNo string `json:"shopee_policy_no"`
	PolicyNo       string `json:"policy_no"`
	LapseTime      string `json:"lapse_time"`
	LapseType      string `json:"lapse_type"`
	Message        string `json:"message"`
}

func main() {
	underWritingResultUrl	 := ShopeeMYClaimHost + "/api/broker/notify_policy_lapse"
	client := &http.Client{Transport: SigTripper{http.DefaultTransport}}
	//req := &UnderwritingParam{
	//	ShopeePolicyNo:    "1534492665016483846_1531207257298044928",
	//	PolicyNo:          "PSP20210800028",
	//	PolicyContractUrl: "https://api.qa.iglooinsure.com/v1/admin/fileapi/common_file/proxy/shopeemy/U2hvcGVlTVktUFNQLTIwMjEtMDgtYmEwYmI2LVBTUDIwMjEwODAwMDI4XzYxMGJiNTc4LnBkZg==",
	//	StartTime:         "2021-08-05T09:54:25+08:00",
	//	EndTime:           "2022-08-05T09:54:25+08:00",
	//	Status:            "surrender",
	//	Message:           "",
	//	ClaimUrl:          "https://gg-shopeemy-psp.qa.axinan.com/claimFlow?id=PSP20210800028",
	//}
	req := &PolicyLapseParam{
		ShopeePolicyNo: "1534492665016483846_1531207257298044928",
		PolicyNo:       "PSP20210800028",
		LapseTime:      time.Now().Format(time.RFC3339),
		LapseType:      "expiration",
		Message:        "",
	}
	data, err := json.Marshal(req)
	if err != nil {
		return
	}

	resp, err := client.Post(underWritingResultUrl,"application/json",bytes.NewReader(data) )
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}

func main1() {
	claimUrl := ShopeeMYClaimHost + "/api/broker/notify_claim_info_v2"
	client := &http.Client{Transport: SigTripper{http.DefaultTransport}}
	//req := &NotifyClaimInfoReq{
	//	ShopeePolicyNO:      "xtz",
	//	ClaimNO:             "fsa",
	//	PolicyNO:            "fff",
	//	ActionTime:          "2021-07-08T15:19:22Z",
	//	ClaimAmount:         "200.00",
	//	ClaimAmountApproved: "200.00",
	//	Status:              "initiated",
	//	Message:             "xxx",
	//}
	req2 := &NotifyClaimInfoV2{
		ClaimNotificationNo: "xxx",
		LossOccurredType:    "damage",
		LossOccurredTime:    time.Now().Format(time.RFC3339),
		Policies:            []*RefPolicyInfo{
			{
				ShopeePolicyNo: "124_1503949777610473472",
				PolicyNo:       "policy_no_002",
				Currency:       "MYR",
			},
		},
		ActionTime:          time.Now().Format(time.RFC3339),
		Status:              "initiated",
	}
	data, err := json.Marshal(req2)
	if err != nil {
		return
	}

	resp, err := client.Post(claimUrl,"application/json",bytes.NewReader(data) )
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
	//r, err  := http.NewRequest(http.MethodPost, claimUrl, bytes.NewReader(data))
	//if err != nil {
	//	return
	//}
	//r.Header.Set("X-Req-AppId", ShopeeMYReqAppId)
	//timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	//r.Header.Set("X-Req-Timestamp", timestampStr)
	//nonceStr, err := GenerateRandomString(15)
	//if err != nil {
	//	return
	//}
	//r.Header.Set("X-Req-Nonce", nonceStr)
	//method := r.Method
	//url := r.URL.Path
	//data, err := json.Marshal(req2)
	//if err != nil {
	//	return
	//}
	//sig, err := GenerateSignature(AppSecretKey, method, url, timestampStr, nonceStr, string(data))
	//if err != nil {
	//	return
	//}
	//r.Header.Set("X-Req-Signature", sig)
	//
	//client1 := &http.Client{}
	//do, err := client1.Do(r)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(do)
	//body, err := ioutil.ReadAll(do.Body)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(body))
}

//package main
//
//import (
//	"fmt"
//	"net/http"
//)
//
//// This type implements the http.RoundTripper interface
//type LoggingRoundTripper struct {
//	Proxies http.RoundTripper
//}
//
//func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
//	// Do "before sending requests" actions here.
//	fmt.Printf("Sending request to %v\n", req.URL)
//
//	// Send the request, get the response (or the error)
//	res, e = lrt.Proxies.RoundTrip(req)
//
//	// Handle the result.
//	if (e != nil) {
//		fmt.Printf("Error: %v", e)
//	} else {
//		fmt.Printf("Received %v response\n", res.Status)
//	}
//
//	return
//}
//
//func main() {
//	httpClient := &http.Client{
//		Transport: LoggingRoundTripper{http.DefaultTransport},
//	}
//	httpClient.Get("https://www.google.com/")
//}