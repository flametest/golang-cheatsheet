package main

import (
    "fmt"
    "strconv"
    "time"
)

//func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
//    var s V
//    for _, v := range m {
//        s += v
//    }
//    return s
//}

func main() {
    for i := 0; i < 100; i++ {
        s := fmt.Sprintf("PL%06s" , strconv.FormatInt(int64(i+1), 10))
        fmt.Println(s)
    }
    fmt.Println(time.Now().Unix())
}
