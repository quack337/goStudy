package main

import (
	"fmt"
	"strconv"
)

func sum(values... any) (result int) {
	for _, value := range values {
		switch value.(type) {
		case int: 
			result += value.(int)
		case string: 
			temp, _ := strconv.Atoi(value.(string))
			result += temp
		}
	}
	return result
}

func main() {
	fmt.Println(sum())
	fmt.Println(1)
	fmt.Println(sum("2", 3))
	fmt.Println(sum(3, "4", 5))
}