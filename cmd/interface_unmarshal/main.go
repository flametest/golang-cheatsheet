package main

import (
	"encoding/json"
	"fmt"
)

type AS struct {
	A string `json:"a"`
}

type BS struct {
	B string `json:"b"`
}

type IS interface {
	Marshal() ([]byte, error)
}

func (a *AS) Marshal() ([]byte, error) {
	return json.Marshal(a)
}

func (b *BS) Marshal() ([]byte, error) {
	return json.Marshal(b)
}

type Unpack struct {
	IS
}

func (u *Unpack) Unmarshal(bytes []byte) error  {
	typeBS := &BS{}
	err := json.Unmarshal(bytes, typeBS)
	if err == nil && u.IS != nil {
		u.IS = typeBS
		return err
	}
	fmt.Println(err)
	if _, ok := err.(*json.UnmarshalTypeError); err != nil &&!ok {
		return err
	}
	return nil
}

func main() {
	a := &AS{A:"1"}
	fmt.Println(a)
	c, err:= a.Marshal()
	if err != nil {
		return
	}
	fmt.Println(string(c))
	up := Unpack{}
	err = up.Unmarshal(c)
	if err != nil {
		return
	}
	fmt.Println(up.IS)


	d, err := json.Marshal(nil)
	if err != nil {
		return
	}
	fmt.Println(string(d))
	e := &AS{}
	err = json.Unmarshal(d, e)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e)

}