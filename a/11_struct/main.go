package main

import (
	"fmt"
	"unsafe"
)

type Point2 struct {
  x int
  y int
}

func main() {
  var p1 = Point2{11, 22}
  fmt.Println(p1)

  type Point3 struct {
	  p Point2
	  z int
  }

  var p2 = Point3{Point2{11, 22}, 33}
  fmt.Println(p2)
  fmt.Println(p2.p.x, p2.p, p2.z)

  type Point33 struct {
	  Point2
	  z int
  }

  var p3 = Point33{Point2{11, 22}, 33}
  fmt.Println(p3)
  fmt.Println(p3.x, p3.y, p3.z)
  fmt.Println(p3.Point2.x, p3.Point2.y, p3.z)

  fmt.Println(unsafe.Sizeof(p1), unsafe.Sizeof(p2))

  f(p1) // call by value
  fmt.Println(p1)

  ////////////////
  type _Point struct { x, y int } // private type

  var p4 _Point = _Point{123, 456}
  var p5 *_Point = &_Point{234, 4567}
  fmt.Println(p4, p5)
}

func f(p Point2) {
  fmt.Println(p)
  p.x = 123
  fmt.Println(p)
}