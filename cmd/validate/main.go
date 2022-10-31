package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)


type Req struct {
	Email string `json:"email" validate:"email"`
}

func main() {
	v := validator.New()
	req := &Req{
		Email: "christia..pusong2@gmail.com",
	}
	err := v.Struct(req)
	if err != nil {
		fmt.Println(err)
	}
}
