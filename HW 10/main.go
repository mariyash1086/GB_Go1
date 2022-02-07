package main

import (
	"HW10/fibonachi"
	"flag"
	"fmt"
)

var x = flag.Float64("x", 0, "x")

func main() {
	flag.Parse()
	result, err := fibonachi.FibonachiProcess(*x)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(result)
}
