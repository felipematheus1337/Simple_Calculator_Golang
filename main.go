package main

import "fmt"

func main() {

	var numero1 int64
	var numero2 int64
	var op int32

	fmt.Println("Welcome to the Golang Pratice Calculator, you need to choose two numbers to use")
	fmt.Println("Type the first number: ")
	fmt.Scan(&numero1)
	fmt.Println("Type the second number: ")
	fmt.Scan(&numero2)

	fmt.Printf("\n Numbers chosen : %d %d  \n", numero1, numero2)

	fmt.Println("\n Now you need to choose some options to implement")
	fmt.Println("\n--------------------------------------------------------------------------------")
	fmt.Println("\n--------------------------------------------------------------------------------")

	for ok := true; ok; ok = !false {
		fmt.Println("Press 1 to Addition, Press 2 to Subtraction, Press 3 to Multiplication, Press 4 Division, and Press 0 to leave the Calculator")
		fmt.Scan(&op)
		toLoop := Calc(op, numero1, numero2)
		if toLoop == 0 {
			ok = false
			return
		}
	}

}

func Calc(Operation int32, firstNumber int64, secondNumber int64) int64 {
	switch Operation {

	case 1:
		{
			soma := firstNumber + secondNumber
			fmt.Printf("Your Addition Result is : %d \n", soma)
			return 1
		}

	case 2:
		{
			subtract := firstNumber - secondNumber
			fmt.Printf("Your Subtract Result is : %d \n", subtract)
			return 2
		}

	case 3:
		{
			multipl := firstNumber * secondNumber
			fmt.Printf("Your Multiplication Result is : %d \n", multipl)
			return 3
		}

	case 4:
		{
			divis := firstNumber / secondNumber
			fmt.Printf("Your Division Result is : %d \n", divis)
			return 4
		}

	case 0:
		{
			fmt.Printf("Exiting... the calculator \n")
			return 0
		}

	default:
		{
			fmt.Println("Wrong number-key pressed")
			return -1
		}

	}

}
