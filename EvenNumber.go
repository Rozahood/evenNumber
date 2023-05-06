package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var userInput int
	fmt.Scanln(&userInput)
	x := userInput * 3
	//fmt.Println("double of ", userInput, "is: ", x)

	fmt.Println("Enter a number: ")
	var userInput1 int
	fmt.Scanln(&userInput1)
	y := userInput1 * 2
	//fmt.Println("triple of ", userInput1, "is: ", y)

	oddCounter := y % x
	result := oddCounter % 3
	fmt.Println("remainder of is: ", result)
	//if oddCounter%2 == 0 {
	//	answer := oddCounter % 3
	//	fmt.Println(answer)
	//} else {
	//	fmt.Println(oddCounter)
	//}
}
