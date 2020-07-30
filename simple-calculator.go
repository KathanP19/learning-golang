package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter What would you like to do?\n\n1.Addition\n2.Subtraction\n3.Multiplication\n4.Division\n")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the Option : ")
	scanner.Scan()
	i, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	if i == 1 {
		fmt.Println("You Have Selected to do Addition!\n")
		fmt.Println("Please Enter how many numbers you want to Add: ")
		scanner.Scan()
		j, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		var l int64
		var p int64
		for k := int64(1); k <= j; k++ {
			fmt.Scan(&l)
			p = int64(p) + int64(l)
		}
		fmt.Printf("Answer : %d", p)
	} else if i == 2 {
		fmt.Println("You Have Selected to do Subtraction!\n")
		fmt.Println("Please Enter how many numbers you want to Sub: ")
		scanner.Scan()
		j, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		var l int64
		var p int64
		for k := int64(1); k <= j; k++ {
			fmt.Scan(&l)
			p = int64(p) - int64(l)
		}
		fmt.Printf("Answer : %d", p)
	} else if i == 3 {
		fmt.Println("You Have Selected to do Multiplication!\n")
		fmt.Println("Please Enter how many numbers you want to Mul: ")
		scanner.Scan()
		j, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		var l int64
		var p int64 = 1
		for k := int64(1); k <= j; k++ {
			fmt.Scan(&l)
			p = int64(p) * int64(l)
		}
		fmt.Printf("Answer : %d", p)
	} else if i == 4 {
		fmt.Println("You Have Selected to do Division!\n")
		fmt.Println("Please Enter two numbers you want to Div: ")
		var l float64
		var p float64
		fmt.Scan(&l)
		fmt.Scan(&p)
		d := float64(l) / float64(p)
		fmt.Printf("Answer : %f\n", d)
	}
}
