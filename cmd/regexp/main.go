package main

import (
	"fmt"
	"regexp"
)

func main() {
	res := "9:44\nк\nS\n|||\nActivation\nPlease use another phone to scan\nthe QR Code below\nXAID\nYou will take a photo from the other phone\nafter QR code scanned.\n3712\nb\nVoi)\n65 s\nLTE1.34%.\nO\nT"
	timePattern := "\\s(?P<second>(\\d)+)\\ss\\s"
	re := regexp.MustCompile(timePattern)
	match := re.FindStringSubmatch(res)
	fmt.Println(match)
	fmt.Println(re.SubexpNames())
	result := match[re.SubexpIndex("second")]
	fmt.Println(result)
	//result := compile.FindString(res.Data)
	//fmt.Println("result")
	//fmt.Println(string(strings.TrimSpace(result)))

	vnPhonePattern := "\\((0|\\+84)\\)\\d+"
	re = regexp.MustCompile(vnPhonePattern)
	fmt.Println(re.MatchString("(0)983280987"))
	fmt.Println(re.MatchString("(+84)983280987"))
}
