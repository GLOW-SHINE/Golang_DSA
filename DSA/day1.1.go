package main

import "fmt"

func main() {
	var (
		name      string
		age       int
		lucky_num float64
	)

	fmt.Print("Enter Youe Name:- ")
	fmt.Scanln(&name)

	fmt.Print("Enter Your Age:- ")
	fmt.Scanln(&age)

	fmt.Print("Enter Your Lucky Number:- ")
	fmt.Scanln(&lucky_num)

	fmt.Println("\n-- OUTPUT--")
	fmt.Printf("Hello %s \n", name)
	fmt.Printf("Your Age %d \n", age)
	fmt.Printf("Your Lucky Number %.2f \n", lucky_num)

}
