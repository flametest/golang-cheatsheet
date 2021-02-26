package main

import("github.com/kellydunn/golang-geo"
"fmt")

func main() {
	// Make a few points
	p := geo.NewPoint(30, 120)
	p2 := geo.NewPoint(20, 120)

	// find the great circle distance between them
	dist := p.GreatCircleDistance(p2)
	fmt.Printf("great circle distance: %d\n", dist)
}
