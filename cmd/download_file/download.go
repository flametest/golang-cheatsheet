package main

import (
	"fmt"
	"regexp"
)

func main() {
	//case1 := "https://s3.ap-southeast-1.amazonaws.com/iglooinsure.prod-fileapi/igloo_lite_sdk-2022-08-ea2030-16600144993346709869704319298565.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFIPVJAQNP%2F20220817%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20220817T073650Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=dd219771f557755d2469a487b7c285c74138f3bb152b1802b018e708cafe0cb6"
	//case2 := "https://s3.ap-southeast-1.amazonaws.com/iglooinsure.prod-fileapi/igloo_lite_sdk-2022-08-6d420a-image.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFIPVJAQNP%2F20220809%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20220809T014443Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=d99f12bf63088c477425bd36d87d758c1cb27701506fbbeba74ebec1cbddfbb2"
	//case3 := "https://s3.ap-southeast-1.amazonaws.com/iglooinsure.prod-fileapi/igloo_lite_sdk-2022-08-fd62af-image.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFIPVJAQNP%2F20220817%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20220817T030341Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=35d34c6ee526d74744f8ee711b014b823d5b31a8e8c70ec1f58d7d6ba9f8ceda"

	//case4 := "https://s3.ap-southeast-1.amazonaws.com/iglooinsure.qa-fileapi/igloo_lite_sdk-2022-09-76bd3c-image.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFADTKQNVK%2F20220902%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20220902T061048Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=37fa697e527c906edd93a97019a5aafa3d97d2d6163c401a997ffa87a6d94349"
	//case5 :=  "https://s3.ap-southeast-1.amazonaws.com/iglooinsure.qa-fileapi/igloo_lite_sdk-2022-09-3220b4-image.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFADTKQNVK%2F20220907%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20220907T080725Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=b4a24e6a71e8f6693ca081737a92b2aa6c032d3624d7c865b1f07736e20b1dc1"
	//resp, err := http.Get(case5)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//respBody, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return
	//}
	//s := base64.StdEncoding.EncodeToString(respBody)
	////fmt.Println(s)
	//req := struct {
	//	Image string `json:"image"`
	//	Type string `json:"type"`
	//}{
	//	Image: s,
	//	Type: "base64",
	//}
	//marshal, err := json.Marshal(req)
	//if err != nil {
	//	return
	//}
	//resp1, err := http.Post("https://api.dev.iglooinsure.com/google-vision/", "application/json", bytes.NewBuffer(marshal) )
	//if err != nil {
	//	panic(err)
	//}
	//defer resp1.Body.Close()
	//respBody1, err := ioutil.ReadAll(resp1.Body)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(respBody1))
	//res := &struct {
	//	Success bool `json:"success"`
	//	Data string `json:"data"`
	//	Message string `json:"message"`
	//}{}
	//err = json.Unmarshal(respBody1, res)
	//if err != nil {
	//	panic(err)
	//}
	result := "|| ° | |\n7\n13:59\nActivation\nPlease use another phone to\nscan the QR Code below\nYou will take a photo from the other\nphone after QR code scanned.\n3817\n426\n0\n100 S"
	//fmt.Println(result1)
	if true {
		//fmt.Println(res.Data)
		timePattern := "after QR code scanned.\\s(?P<activationID>(\\d)+)\\s[\\s\\S]*\\s(?P<second>(\\d)+)\\s[sS]"
		re := regexp.MustCompile(timePattern)
		fmt.Println(	re.MatchString(result))
		match := re.FindStringSubmatch(result)
		fmt.Println(match)
		fmt.Println(re.SubexpNames())

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
