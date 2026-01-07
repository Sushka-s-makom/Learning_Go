package main

import "fmt"

func slices_copy() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

}
