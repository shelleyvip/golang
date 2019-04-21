package main

import (
	"fmt"
	"math"
)

type Point struct {
	x,y float64
}

func (p *Point)Distance(q Point)float64  {
	return math.Hypot(p.x-q.y,p.x-q.y)
}

func Distance(path []Point)float64  {

}

func main() {
    path := []Point{{1,2},{3,4},{5,6}}
    fmt.Println(path)
	//p := Point{6,8}
	//q := Point{1,3}
	//fmt.Println(p.Distance(q))
	
}
