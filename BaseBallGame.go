package main 
import (
	"fmt"
"strconv"
)
func calPoints(operations []string) int {
	stack := []int{}

	for _, op := range operations {

		if op == "+" {
			n := len(stack)
			stack = append(stack, stack[n-1]+stack[n-2])

		} else if op == "D" {
			n := len(stack)
			stack = append(stack, 2*stack[n-1])

		} else if op == "C" {
			stack = stack[:len(stack)-1]

		} else {
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
	}

	sum := 0

	for _, score := range stack {
		sum += score
	}

	return sum
}
func main() {
	result:=calPoints([] string{""})
	fmt.Println(result)
	
}