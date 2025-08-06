package main

import "fmt"

func main() {
	var name string = "PALASH"
	var age int = 30
	var lucky_num float64 = 1

	fmt.Print("Enter your Name : - ")
	fmt.Scanln(&name)
	fmt.Println(name)
	fmt.Print("Enter Your Age")
	fmt.Scanln(&age)
	fmt.Println(age)
	fmt.Print("Enter Lucky Number :- ")
	fmt.Scan(&lucky_num)
	fmt.Println(lucky_num)

	/* fmt.Printf("Hello : %s \n", name)
	fmt.Printf("Age : %d \n", age)
	fmt.Printf("Your lucky number is : %.2f \n", lucky_num)
	*/
}
