package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr = "{\n    \"label_dic\": {\n        \"blur\": 0.0247,\n        \"cover\": 0.0021,\n        \"crack\": 0.0204,\n        \"dark\": 0.0002,\n        \"far\": 0,\n        \"fraud\": 0.0006,\n        \"likely_crack\": 0.0782,\n        \"noncrack\": 0.8737,\n        \"others\": 0\n    },\n    \"url\": \"https://storage.googleapis.com/axinan_psp_prod/croped/2020/09/47b8a/photo.jpg\"\n}"

type ActivationResultType string

const (
	ActivationResultTypeCover       ActivationResultType = "cover"
	ActivationResultTypeCrack       ActivationResultType = "crack"
	ActivationResultTypeFar         ActivationResultType = "far"
	ActivationResultTypeFraud       ActivationResultType = "fraud"
	ActivationResultTypeDark        ActivationResultType = "dark"
	ActivationResultTypeBlur        ActivationResultType = "blur"
	ActivationResultTypeLikelyCrack ActivationResultType = "likely_crack"
	ActivationResultTypeOthers      ActivationResultType = "others"
	ActivationResultTypeNoncrack    ActivationResultType = "noncrack"
)

type ProcessedResult struct {
	LabelDic map[ActivationResultType]float64 `json:"label_dic"`
	Url      string `json:"url"`
}


func main() {
	pr := &ProcessedResult{}
	err := json.Unmarshal([]byte(jsonStr), pr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pr)
	for resultType, _ := range pr.LabelDic {
		fmt.Println(resultType == ActivationResultTypeOthers)
	}
	jstr, err := json.Marshal(pr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jstr))
}
