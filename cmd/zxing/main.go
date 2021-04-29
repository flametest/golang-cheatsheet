package main

import (
	"fmt"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"image"
	_ "image/jpeg"
	"os"
)

func main() {
	file, err := os.Open("cmd/zxing/qr_code.jpg")
	if err != nil {
		fmt.Println(err)
	}
	img, _, err := image.Decode(file)
	//resp, err := http.Get("https://s3.ap-southeast-1.amazonaws.com/iglooinsure.qa-fileapi/igloo_lite_sdk-2021-04-c667e0-image.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAU2NHV2XFADTKQNVK%2F20210428%2Fap-southeast-1%2Fs3%2Faws4_request&X-Amz-Date=20210428T105834Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=b6cab8eebfbaee206a17c6099cdf7ff31840c7f1946c63d9227d84056022effe")
	//if err != nil {
	//	return
	//}
	//img, _, err := image.Decode(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)
	qrReader := qrcode.NewQRCodeReader()
	result, _ := qrReader.Decode(bmp, nil)

	fmt.Println(result)
}
