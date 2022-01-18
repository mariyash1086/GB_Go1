package main

import "fmt"

func main() {
	var n uint
	fmt.Scanln(&n)
	fmt.Println("Фибоначчи числа равен ", fibonachi(n))
}

func fibonachi(n uint) uint {

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibonachi(n-1) + fibonachi(n-2)

}
