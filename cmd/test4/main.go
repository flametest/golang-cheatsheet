package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	id uint64 `json:"id"`
	name string `json:"name"`
}


func worker(id uint64, c chan *Person) {
	time.Sleep(5 * time.Second)
	if id % 2 == 0 {
		c <- &Person{id: id, name: string(rand.Int())}
	} else {
		c <- nil
	}

}

func main() {
	start := time.Now()
	retMap := make(map[uint64]*Person)
	ids := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	channels := make([]chan *Person,len(ids))
	for i, id := range ids {
		channels[i] = make(chan *Person, 1)
		go worker(id, channels[i])
	}
	for i, id := range ids {
		retMap[id] = <-channels[i]
		close(channels[i])
	}
	fmt.Println(retMap)
	fmt.Printf(time.Now().Sub(start).String())
}
