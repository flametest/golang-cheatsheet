package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type Tripper struct {
	next http.RoundTripper
}

func NewTripper(n http.RoundTripper) *Tripper {
	return &Tripper{next :n}
}

func (t *Tripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	fmt.Println("xxx")
	return t.next.RoundTrip(req)
}

func main() {
	//trans := http.DefaultTransport
	//tripper := NewTripper(trans)
	//client := &http.Client{Transport: tripper}
	//client.Get("www.google.com")
	rawPayload, err := base64.RawURLEncoding.DecodeString("eyJlbWFpbCI6Im5ld3RvbmRAYXhpbmFuLmNvbSIsImV4cCI6MTYyMTA2NzA3MCwidWlkIjo2fQ")
	if err != nil {
fmt.Println(err)
	}
	fmt.
		Println(rawPayload)
}
