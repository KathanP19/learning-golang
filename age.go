package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	year := int64(t.Year())
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type Your Name : ")
	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("Hello %s Lets Calculate your Age\n\nEnter the Year your were Born In : ", input)
	scanner.Scan()
	newInput, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("Your Age is %d\n", year-newInput)
}
