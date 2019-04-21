package main

import (
	"math"
)

type Point struct {
	x,y float64
}

func (p Point)Distance2Point(q Point)float64  {
	return math.Hypot(p.x-q.x,p.y-q.y)
}

func (p Point)Distance()float64  {
	return math.Hypot(p.x,p.y)
}

type Path []byte

func (p Path)Distance() float64{
  var sum float64
  for i := 0;i <len(p)-1;i++{
  	sum += p[i].
  }
  return sum
}

func main() {
	
}
