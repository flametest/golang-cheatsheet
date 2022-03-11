package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"strings"
)

const AESKey = "3f22c60a-c826-11eb-a05a-b8599f49e654"

func DecryptUserId(encryptedUserId string) (string, error) {
	base64Decode, err := base64.StdEncoding.DecodeString(encryptedUserId)
	if err != nil {
		return "", err
	}
	result, err := AESDecrypt(base64Decode, []byte(AESKey)[:aes.BlockSize])
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func AESDecrypt(crypt []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(crypt) == 0 {
		return nil, fmt.Errorf("plain content empty")
	}
	iv := make([]byte, aes.BlockSize)
	ecb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return pksc5Trimming(decrypted), nil
}

func pksc5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func main() {
	fmt.Println(DecryptUserId("AoqzxAve7146Bl1/hzXn/sKQFj7/tzWJ5G6KYnenJao="))
	phoneStrings := strings.Split("63-9686503496", "-")
	fmt.Println(phoneStrings[0], phoneStrings[1])
}
