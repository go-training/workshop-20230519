package main

import "fmt"

func hello() string {
	return "歡迎來到 Go 世界"
}

func foobar(name string) string {
	return "Welcome to " + name
}

func main() {
	fmt.Println(hello())
	fmt.Println(foobar("foobar"))
}
