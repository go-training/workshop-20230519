//go:build !go1.20
// +build !go1.20

package hello3

import (
	"runtime"
)

// Hello is hello world
func Hello() string {
	return "build on " + runtime.Version()
}
