package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {		
	var reader = bufio.NewReader(os.Stdin)
	var s,_ = reader.ReadString('\n')
	fmt.Println(s)
}