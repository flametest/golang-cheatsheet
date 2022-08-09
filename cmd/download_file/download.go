package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	//case1 := "https://s3.ap-southeast-1.amazonaws.com/iglooinsure.prod-fileapi/igloo_lite_sdk-2022-08-ea2030-16600144993346709869704319298565.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFIPVJAQNP%2F20220817%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20220817T073650Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=dd219771f557755d2469a487b7c285c74138f3bb152b1802b018e708cafe0cb6"
	//case2 := "https://s3.ap-southeast-1.amazonaws.com/iglooinsure.prod-fileapi/igloo_lite_sdk-2022-08-6d420a-image.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFIPVJAQNP%2F20220809%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20220809T014443Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=d99f12bf63088c477425bd36d87d758c1cb27701506fbbeba74ebec1cbddfbb2"
	case3 := "https://s3.ap-southeast-1.amazonaws.com/iglooinsure.prod-fileapi/igloo_lite_sdk-2022-08-fd62af-image.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFIPVJAQNP%2F20220817%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20220817T030341Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=35d34c6ee526d74744f8ee711b014b823d5b31a8e8c70ec1f58d7d6ba9f8ceda"
	resp, err := http.Get(case3)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	s := base64.StdEncoding.EncodeToString(respBody)
	//fmt.Println(s)
	req := struct {
		Image string `json:"image"`
		Type string `json:"type"`
	}{
		Image: s,
		Type: "base64",
	}
	marshal, err := json.Marshal(req)
	if err != nil {
		return
	}
	resp1, err := http.Post("https://api.dev.iglooinsure.com/google-vision/", "application/json", bytes.NewBuffer(marshal) )
	if err != nil {
		panic(err)
	}
	defer resp1.Body.Close()
	respBody1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody1))
	res := &struct {
		Success bool `json:"success"`
		Data string `json:"data"`
		Message string `json:"message"`
	}{}
	err = json.Unmarshal(respBody1, res)
	if err != nil {
		panic(err)
	}
	if res.Success {
		//fmt.Println(res.Data)
		timePattern := "after QR code scanned.\\s(?P<activationID>(\\d)+)\\s.*\\s(?P<second>(\\d)+)\\ss\\s"
		re := regexp.MustCompile(timePattern)
		match := re.FindStringSubmatch(res.Data)
		//fmt.Println(match)
		//fmt.Println(re.SubexpNames())
		if re.SubexpIndex("activationID") != -1 {
			result := match[re.SubexpIndex("activationID")]
			fmt.Println(result)
		}
		if re.SubexpIndex("second") != -1 {
			result := match[re.SubexpIndex("second")]
			fmt.Println(result)
		}

	}
}
