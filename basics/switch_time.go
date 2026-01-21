package main

import (
	"fmt"
	"time"
)

func switch_time() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}
}
