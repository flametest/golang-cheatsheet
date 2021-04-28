package main

import (
	"errors"
	"fmt"
	"github.com/speps/go-hashids"
)

var (
	hid, _ = hashids.NewWithData(
		&hashids.HashIDData{
			Salt: "HRceEfXKZ3lDqdjbRc3am/TMqS55ZQ4NveZM87HgIoE=",
			// so that we have a minimum length of 28 and each with 26+10 combination
			// should be safe against hash collision
			MinLength: 15,
			Alphabet:  "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		},
	)
)

type IDConverter struct {
	idPrefix string
}

func NewIDConverter(idPrefix string) *IDConverter {
	return &IDConverter{idPrefix: idPrefix}
}

func (p *IDConverter) ToString(id int64) (string, error) {
	hashid, err := hid.EncodeInt64([]int64{id})
	if err != nil {
		return "", err
	}
	return p.idPrefix + hashid, err
}
func (p *IDConverter) ToInt64(id string) (int64, error) {
	hashid := id[len(p.idPrefix):]
	is, err := hid.DecodeInt64WithError(hashid)
	if err != nil {
		return 0, err
	}
	if len(is) != 1 {
		return 0, errors.New("decode failed")
	}
	return is[0], nil
}

func main() {
	c := NewIDConverter("PSP")
	s, err := c.ToString(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	i, err := c.ToInt64(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}

