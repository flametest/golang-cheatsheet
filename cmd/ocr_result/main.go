package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	values := url.Values{}
	values.Set("image_url", "")
	resp, err := http.PostForm("https://aiplatform.dev.axinan.com/port5/ ", values)
	if err != nil {
		panic(err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))
	pr := &OCRResult{}
	err = json.Unmarshal(respBody, pr)
	if err != nil {
		panic(err)
	}
	fmt.Println(pr)
	fmt.Println(time.Now().Format(time.RFC3339))
}
