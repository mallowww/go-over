package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	/*
		if err != nil {
			msg, _ := fmt.Scanf("%v must be a number only", promt)
			panic(msg)
		}
	*/
	if err != nil {
		input = strings.TrimRight(input, "\r\n")
		message, err := fmt.Printf("%s ERR: not string must number only \n", input)
		if err != nil {
			fmt.Println(err)
		}
		panic(message)
	}
	return value
}

func getOperator() string {
	fmt.Println("operator is ( + - * / ): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func sum(a, b float64) float64 {
	return a + b
}

func int_sum(xs ...int) int{
	int_sum := 0
	for _, x := range xs {
		int_sum += x
	}
	return int_sum
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}
func divide(a, b float64) float64 {
	return a / b
}

func main() {
	a := getInput(" value1 = ")
	b := getInput(" value2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = sum(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		panic("wrong operator")
	}
	fmt.Printf("result = %v", result)

}
