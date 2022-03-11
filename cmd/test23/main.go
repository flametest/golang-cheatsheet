package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type S1 struct {
	A string
}

func (s S1) E()  {

}

type S2 struct {
	S1
}

type IS interface {
	E()
}



func main() {
	fmt.Println(strings.HasPrefix("+63", "+", ))
	fmt.Println(time.Now().Format("2006-01-02"))
	app := fmt.Sprintf("%v", "gadget")
	fmt.Println(app)
	var LocationBangkok, _ = time.LoadLocation("Asia/Bangkok")
	start := time.Date(2021, 12,02, 0,0,0,0, LocationBangkok)
	fmt.Println(start.Unix())
	end := time.Now().Add(time.Hour * 24 * 180)
	fmt.Println(end.Format(time.RFC3339))
	fmt.Println(end.Unix())
	var a IS
	a = S2{}
	aType := reflect.TypeOf(a)
	fmt.Println(aType)
	f, ok := aType.FieldByName("S1")
	fmt.Println(f.Name, ok)
}
