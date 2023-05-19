package main

import "fmt"

func main() {
	foo01()
	foo02()
	foo03()
	foo04()

	foo := []int{1, 2, 3, 4}
	calc(1, 2, 3)
	calc(foo...)
	calc(1, 2, 3, 4)
}

func foo01() {
	var foo [3]int  // array
	var bar []int   // slice
	test := []int{} // slice

	if bar == nil {
		fmt.Printf("%v\n", foo)
		fmt.Printf("%p\n", &foo)
		fmt.Printf("%p\n", bar)
		fmt.Printf("%p\n", test)
	}
}

func foo02() {
	x := []int{1, 2, 3}
	y := x[1:]          // [2, 3]
	y[0] = 100          // y = [100, 3]
	fmt.Println(x)      // [1, 100, 3]
	fmt.Println(y)      // [100, 3]
	y = append(y, 1000) // y = [100, 3, 1000]
	y[0] = 99           // y = [99, 3, 1000]
	fmt.Println(x)      // [1, 100, 3]
	fmt.Println(y)      // [99, 3, 1000]
}

func foo03() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) // [7, 2, 3]
	}(x)
	fmt.Println(x) // [1, 2, 3]
}

func foo04() {
	x := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(x) // [7, 2, 3]
	}(x)
	fmt.Println(x) // [7, 2, 3]
}

// vals = []int{1, 2, 3}
func calc(vals ...int) {
	foo := 1
	if len(vals) > 0 {
		foo = vals[0]
	}

	bar := append([]int{}, vals...)
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println("array count:", len(vals))
	for index, val := range vals {
		fmt.Println(index, ":", val)
	}
}
