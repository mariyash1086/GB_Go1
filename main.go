package main

import "fmt"

//HW4
func main() {
	slice := []int{1, 10, 20, 5, 6, 4, 0, 178, 150}

	fmt.Println(slice)
	sort(slice)
	fmt.Println(slice)
}

func sort(a []int) {

	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1

		for j := 1; j < len(a); {
			if i >= 0 && a[i] > key {
				a[i+1] = a[i]
				i = i - 1

			} else {
				break

			}
			a[i+1] = key
		}

	}
}
