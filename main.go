package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

//hw3
func main() {

	//1. calculation
	calc()
	//2. prime numbers
	numbers()

}

func calc() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите простую арифметическую операцию (+, -, *, /) или сложную " +
		"(x-возведение А в степень B, y- извлечение квадратного корня из A, z -извлечение корня кубического из A)")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "x":
		res = math.Pow(a, b)
	case "y":
		res = math.Sqrt(a)
	case "z":
		res = math.Pow(a, 1.00/3.00)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)
}

func numbers() {

	var i, n int64

	fmt.Println("Найдем все простые числа от 0 до N. Введите число N")
	fmt.Scanln(&n)
	for i = 0; i <= n; i++ {
		if i <= 1 {
			fmt.Println(i, "-Это НЕ простое число")
			continue
		}
		if big.NewInt(i).ProbablyPrime(0) {
			fmt.Println(i, "-Это простое число")
		} else {
			fmt.Println(i, "-Это НЕ простое число")
		}

	}

}
