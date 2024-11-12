package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	double := createTrasfomer(2)
	triple := createTrasfomer(3)
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	fmt.Println(palin(10))
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func createTrasfomer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}

func palin(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * palin(num-1)
	}
}
