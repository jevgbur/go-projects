package main
import "fmt"

func main() {
	var a, b int
	var message string

	fmt.Println("Enter expression:")
	fmt.Scanf("%d %s %d", &a, &message, &b) // Read input in the format: number_operator_number

	switch message {
	case "-":
		fmt.Println(a - b)
	case "+":
		fmt.Println(a + b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero!")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Unknown operation")
	}
}
