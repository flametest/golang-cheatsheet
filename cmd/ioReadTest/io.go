package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func IOReadAll(i io.Reader) {
	_, err := ioutil.ReadAll(i)
	if err != nil {
		fmt.Errorf("readall err:%v", err)
		return
	}
}

func IOCopy(i io.Reader) {
	data := make([]byte, 0)
	buf := bytes.NewBuffer(data)
	_, err := io.Copy(buf, i)
	if err != nil {
		fmt.Errorf("readall err:%v", err)
		return
	}
}
