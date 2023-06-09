package main

import (
	"fmt"
)

type ipv4 []byte

// Stringer is implemented by any value that has a String method,
// which defines the “native” format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
func (s ipv4) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", s[0], s[1], s[2], s[3])
}

// GoString method, which defines the Go syntax for that value.
// The GoString method is used to print values passed as an operand
// to a %#v format.
func (s ipv4) GoString() string {
	return fmt.Sprintf("%v.%v.%v.%v", s[0], s[1], s[2], s[3])
}

func main() {
	ips := map[string]ipv4{
		"localhost": {127, 0, 0, 1},
		"Google":    {8, 8, 8, 8},
	}

	for i, v := range ips {
		fmt.Printf("name: %s, ip: %v\n", i, v)
		fmt.Printf("debug: %#v\n", v)
	}
}
