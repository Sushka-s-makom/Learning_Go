package main

import "fmt"

func demoVariables() {
	var a string = "intial"
	fmt.Println(a)

	var b, c int = 11, 18
	fmt.Println(b, c)

	var p int
	fmt.Println(p)

	// сокращения для объявления инициализированной переменной
	f := "short"
	fmt.Println(f)
}
