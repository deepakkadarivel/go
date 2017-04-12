package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg int
	prob string
}

func (e argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, argError{arg, "Can't work with it."}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{42, 7} {
		f1(i)
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed : ", e)
		} else {
			fmt.Println("f1 worked : ", r)
		}
	}

	for _, i := range []int{42, 30} {
		f2(i)
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed : ", e)
		} else {
			fmt.Println("f3 worked : ", r)
		}
	}
}
