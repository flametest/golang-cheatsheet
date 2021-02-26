package main

import (
	"fmt"
)
import "github.com/go-playground/validator/v10"

type ActivatePolicyRequest struct {

	// activation form
	ActivationForm *ActivationForm `json:"activation_form,omitempty" validate:"required"`

	Cost int `json:"cost" validate:"required,numeric,min=13,max=13"`
}

type ActivationForm struct {

	// insured object
	InsuredObject *InsuredObject `json:"insured_object,omitempty" validate:"required"`

	// insured person
	InsuredPerson *InsuredPerson `json:"insured_person,omitempty" validate:"required"`
}

type InsuredPerson struct {

	// birthdate
	Birthdate string `json:"birthdate,omitempty" validate:"required,containsany=ee@yy"`

	// identity photo
	IdentityPhoto []string `json:"identity_photo" validate:"required,gt=0,dive,url"`
}

type InsuredObject struct {

	// i m e i number
	IMEINumber string `json:"IMEI_number,omitempty"`

	// i m e i photo
	IMEIPhoto string `json:"IMEI_photo,omitempty"`
}

func main() {
	fmt.Println("hello")
	ar := &ActivatePolicyRequest{
		ActivationForm: &ActivationForm{
			InsuredPerson: &InsuredPerson{Birthdate: "ee", IdentityPhoto: []string{"http://xx"}},
			InsuredObject: &InsuredObject{IMEINumber: "11", IMEIPhoto: "11"},
		},
			Cost: 12,
	}
	validate := validator.New()
	err := validate.Struct(ar)
	if err != nil {
		fmt.Println(err.Error())
	}
}
