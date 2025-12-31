// в Go нет тернарного оператора, используем full version operator if/else

package main

import "fmt"

func if_1() {

	n := 100
	if n%2 == 0 {
		fmt.Println("n четное число")
	} else {
		fmt.Println("n нечетное число")
	}

}
