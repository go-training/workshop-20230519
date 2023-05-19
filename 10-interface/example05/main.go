package main

import "fmt"

type StatusType int

const (
	StatusPrepare StatusType = iota // 0
	StatusInitial                   // 1
	StatusRunning                   // 2
	StatusSuccess                   // 3
	StatusFailure                   // 4
)

func main() {
	fmt.Println(
		StatusInitial,
		StatusRunning,
		StatusSuccess,
		StatusFailure,
	) // 1, 2, 3, 4
}
