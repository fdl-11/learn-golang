package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
        total += number
    }

	return total
}

func main() {
	total := sumAll(10, 90, 50, 100)
	fmt.Println(total)

	slice := []int{100, 20, 310}
	total = sumAll(slice...)
	fmt.Println(total)
}