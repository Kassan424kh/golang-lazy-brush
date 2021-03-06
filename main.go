package main

import (
	"fmt"

	LB "github.com/Kassan424kh/golang-lazy-brush/lazy-brush"
	P "github.com/Kassan424kh/golang-lazy-brush/point"
)

func main() {
	lb := LB.Init()

	lb.Update(P.Point{X: 50, Y: 0}, false)
	fmt.Println(lb.GetBrushCoordinates())

	lb.Update(P.Point{X: 200, Y: 50}, false)
	fmt.Println(lb.GetBrushCoordinates())
}
