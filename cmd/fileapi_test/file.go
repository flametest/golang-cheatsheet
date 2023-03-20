package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"
)

func Upload(data[]byte, params map[string]string, paramName, fileName string) error {
	start := time.Now().UnixNano()
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fileName)
	if err != nil {
		return err
	}
	part.Write(data)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("time cost %d", time.Now().UnixNano() - start))
	return nil
}

func IOCopy(i io.Reader) []byte {
	data := make([]byte, 0)
	buf := bytes.NewBuffer(data)
	_, err := io.Copy(buf, i)
	if err != nil {
		fmt.Errorf("readall err:%v", err)
		return nil
	}
	return data
}

func main() {
	file, err := os.Open("/Users/jiangjun/go/src/github.com/flametest/golang-cheatsheet/cmd/ioReadTest/test.pdf")
	if err != nil {
		fmt.Errorf("open err:%v", err)
		return
	}
	params := make(map[string]string)
	params["biz_key"] = "gadget"
	params["is_public"] = "false"
	params["with_content"] = "true"

	for {
		Upload(IOCopy(file), params,  "file", "test.pdf")
		time.Sleep(time.Millisecond * 10)
	}
}
