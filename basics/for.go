package main

import "fmt"

func for_1() {

	i := 1
	for i <= 3 {

		fmt.Println(i)
		i += 1
	}

	for j := 7; j < 13; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("plsss, type break")

		break
	}

}
