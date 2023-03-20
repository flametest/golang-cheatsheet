package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
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

func signQueryParam(queryParam map[string][]string, appSecret string) string {
	queryKeyList := make([]string, 0)
	for key, valueList := range queryParam {
		sort.Strings(valueList)
		queryParam[key] = valueList
		queryKeyList = append(queryKeyList, key)
	}
	sort.Strings(queryKeyList)
	sortedKeyPairList := make([]string, 0)
	for _, queryKey := range queryKeyList {
		if queryKey == "sign" {
			continue
		}
		values := queryParam[queryKey]
		for _, value := range values {
			keyPair := queryKey + "=" + value
			sortedKeyPairList = append(sortedKeyPairList, keyPair)
		}

	}
	queryString := strings.Join(sortedKeyPairList, "&")
	s := queryString + appSecret

	sumMD5 := md5.Sum([]byte(s))
	var ourSign string

	ourSign = hex.EncodeToString(sumMD5[:])

	fmt.Printf("our sign is %s", ourSign)

	return ourSign
}

func main() {
	//fmt.Println(QRCodeSignatureV1("test_salt", "GC-2021-8CFED5", "79d7349834315a78", "", 1640165165))
	a := signQueryParam(map[string][]string{
		"timestamp": []string{"1669966861916"},
		"random": []string{"438074"},
	}, "nmxfaozNdCPOsOQ9vvVuQ")
	fmt.Println(a)
}
