package main

import (
	"fmt"
	"github.com/flametest/test/cmd/test24/vo"
)

type CustomerMeta struct {
	CustomerAddress *vo.Address `json:"customer_address"`
	IdentityImages  vo.Photos   `json:"identity_images"`
	Promotions      bool        `json:"promotions"`
}

func main() {
	meta := CustomerMeta{}
	fmt.Println(meta.IdentityImages)
	fmt.Println(meta.CustomerAddress)
}
