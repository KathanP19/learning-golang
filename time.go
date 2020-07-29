package main

import (
	"fmt"
	"time"
)

func main() {

	dt := time.Now()
	fmt.Println(dt.Format("02-01-2006 15:04:05 Monday"))

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
