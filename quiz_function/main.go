package main

import "fmt"

func main() {
	scores := []int{10, 5, 8, 9, 7}
	sum := sum(scores)
	calculation, err := calculate(10, 2, "/")

	fmt.Println("sum : ", sum)
	fmt.Println("calculation : ", calculation, "error : ", err)
}

func sum(scores []int) (sum int) {
	for _, score := range scores {
		sum += score
	}
	return
}

func calculate(num1 int, num2 int, operand string) (hasil int, err string) {
	switch operand {
	case "*":
		hasil = num1 * num2
		return
	case "+":
		hasil = num1 + num2
		return
	case "-":
		hasil = num1 - num2
		return
	case "/":
		hasil = num1 / num2
		return

	default:
		err = "operand error"
		return
	}

}
