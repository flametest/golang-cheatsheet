package main

import (
	"crypto/sha256"
	"fmt"
)

func QRCodeSignatureV1(salt, ID, androidID, UUID string, timestamp int64) string {
	sum := sha256.Sum256([]byte(
		fmt.Sprintf("%s:%s:%s:%d:%s", ID, androidID, UUID, timestamp, salt)),
	)
	return fmt.Sprintf("%x", sum)
}

func PhoneSignature(androidID, uuid string) string {
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s:%s", androidID, uuid)))
	return fmt.Sprintf("%x", sum)
}

func main() {
	fmt.Println(QRCodeSignatureV1("test_salt", "GC-2021-8CFED5", "79d7349834315a78", "", 1640165165))
}
