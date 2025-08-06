package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Your Name :-")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Your AGE :-")
	ageInput, _ := reader.ReadString('\n')
	ageInput = strings.TrimSpace(ageInput)
	age, _ := strconv.Atoi(ageInput)

	fmt.Print("Enter Your Lucky Number :-")
	luckyInput, _ := reader.ReadString('\n')
	luckyInput = strings.TrimSpace(luckyInput)
	luckyNum, _ := strconv.ParseFloat(luckyInput, 64)

	fmt.Println("\n--OUTPUT--")
	fmt.Printf("Hello %s\n", name)
	fmt.Printf("your AGE : %d\n", age)
	fmt.Printf("Your Lucky Number: %.2f\n", luckyNum)

}
