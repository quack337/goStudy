package main

import "fmt"

type Point struct { x, y int }
type Person struct { name string; age int }

func testType(v any) {
	switch t := v.(type) {
	case int: fmt.Println("int", v, t)
	case string: fmt.Println("string", v, t)
	case bool: fmt.Println("bool", v, t)
	case Point: fmt.Println("Point", v, t, v.(Point).x)
	default: fmt.Println("unknown type", v, t)
	}
}

func testTypeCasting(v any) {
	//var person = v.(Person) // runtime error
	var person, ok = v.(Person) // runtime error 아님
	if (ok) {
		fmt.Println(person.name, person.age)
	} else {
		fmt.Println("not person type")
	}
}

func main() {
	testType(3)
	testType("hello")
	testType(true)
	testType(Point{3,4})
	testTypeCasting(Point{3,4})
}