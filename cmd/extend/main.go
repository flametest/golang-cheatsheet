package main

import "fmt"

type IBird interface {
	SayName(bird IBird) string
	Name() string
}

type ABird struct{
}

func (A *ABird) SayName(bird IBird) string {
	return bird.Name()
}

func (A *ABird) Name() string {
	return "a"
}

type BBird struct {
	ABird
}

func (B *BBird) Name() string {
	return "b"
}

type CBird struct {
	ABird
}

func (C *CBird) Name() string {
	return "c"
}

func FactoryCreate() IBird {
	return &CBird{}
}

func SaveBird(bird IBird) {
	fmt.Println(bird.SayName(bird))
}

func main() {
	f := FactoryCreate()
	SaveBird(f)
}

