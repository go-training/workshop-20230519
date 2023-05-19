package main

import (
	"errors"
	"fmt"
)

// return value without define variable.
func hello(a, b int) (int, bool, error) {
	if a > b {
		return 1, true, errors.New("Error: a > b")
	}

	return 0, false, errors.New("Error: a < b")
}

// return value with define variable.
func hello2(a, b int) (val int, check bool, err error) {
	if a > b {
		val = 1
		check = true
		err = errors.New("Error: a > b")
		return
	}

	err = errors.New("Error: a < b")
	return
}

func variable() {
	x := 1
	fmt.Println(x) // 1
	{
		fmt.Println(x) // 1
		x := 2
		fmt.Println(x) // 2
	}
	fmt.Println(x) // 1
	fmt.Println()
}

func main() {
	_, check, _ := hello(1, 2)
	if check {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if _, _, err := hello(1, 2); err != nil {
		fmt.Println("true")
	} else {
		fmt.Println(err)
	}

	if _, _, err := hello(1, 2); err == nil {
		fmt.Println("true")
	} else {
		fmt.Println(err)
	}

	variable()

	// variable scope
	for i := 0; i < 5; i++ {
		defer fmt.Println("defer to:", i)
		defer func() {
			// how to fix the i variable?
			fmt.Println("defer func to:", i)
		}()
	}

	// random order list
	foo := map[string]string{
		"foo01": "test01",
		"bar02": "test02",
		"foo03": "test03",
		"bar04": "test04",
		"foo05": "test01",
		"bar06": "test02",
		"foo07": "test03",
		"bar08": "test04",
		"bar09": "test04",
		"bar10": "test04",
	}

	for i, v := range foo {
		fmt.Println(i, v)
	}

	// read channel
	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	for v := range c {
		fmt.Println("chan value:", v)
	}
}
