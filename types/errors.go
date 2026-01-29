package types

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 0 {
		return -1, errors.New("can't work with 0")
	}
	return arg, nil
}

var ErrOutOfTea = errors.New("no more tea available")

func makeTea(arg int) (int, error) {
	if arg < 2 {
		return 0, ErrOutOfTea
	}
	fmt.Println("making tea)")
	return arg, nil
}
