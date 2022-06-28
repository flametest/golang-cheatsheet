package main

import (
	"fmt"
	"net/url"
)

func main ()  {
	params := url.Values{}
	params.Add("email", "xx@gmail.com")
	baseUrl, err := url.Parse("https://gg-shopeemy.staging.axinan.com")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}
	baseUrl.RawQuery = params.Encode()
	fmt.Println(baseUrl.String())
	params1 := url.Values{}
	params1.Add("phone", "11111")
	baseUrl.RawQuery = params1.Encode()
	fmt.Println(baseUrl.String())
}