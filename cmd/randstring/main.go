package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(n int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GetRandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, length)
	for i := 0; i < length; i++ {
		idx := rand.Intn(len(letters))
		s[i] = rune(letters[idx])
	}
	return string(s)
}

func main() {
	s, err := GenerateRandomString(112)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(GetRandString(6))
	fmt.Println(time.Time{}.Unix())
	var a []*string
	a = nil
	b := "a"
	fmt.Println(a)
	a = append(a, &b)
	fmt.Println(a)
	keyList := []string{"foo", "foo_bar"}
	sort.Strings(keyList)
	fmt.Println(keyList)
	cx := md5.New().Sum([]byte(strconv.FormatInt(time.Now().UnixNano(), 10)))
	fmt.Println(base64.StdEncoding.EncodeToString(cx))
	prefix := "60"
	userPhone := "+60194562610"
	if strings.Contains(userPhone, "+") {
		userPhone = strings.TrimPrefix(userPhone, "+"+prefix)
	}
	fmt.Println(userPhone)
}


