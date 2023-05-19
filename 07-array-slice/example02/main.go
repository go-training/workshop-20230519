package main

import "fmt"

func main() {
	foo()
	bar()
}

func foo() {
	a := [3]int{1, 2, 3}
	for k, v := range a { // copy
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a) // [100, 200, 3]
		}
		a[k] = 100 + v // a[0] = 100 + 1, a[1] = 100 + 2, a[2] = 100 + 3
	}
	fmt.Println(a) // [101, 102, 103]
}

func bar() {
	a := []int{1, 2, 3}
	for k, v := range a { // reference
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a) // [100, 200, 3]
		}
		a[k] = 100 + v // a[0] = 100 + 1, a[1] = 100 + 200, a[2] = 100 + 3
	}
	fmt.Println(a) // [101, 300, 103]
}
