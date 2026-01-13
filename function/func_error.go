package main

import (
	"errors"
)

func f(arg int) (int, error) {

	if arg == 22 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 10, nil
}
