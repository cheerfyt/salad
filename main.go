package main

import (
	"github.com/tate-fan/swgo/swgo"
)

var stack []byte

func charCodeAt(str string, pos int) rune {
	i := 0
	for _, value := range str {
		if i == pos {
			return value
		}
		i++
	}
	return 0
}

func main() {
	debug2 := swgo.NewDebugger("ysp:some")

	for i := 0; i < 2; i++ {
		debug2("shangtian: %s", "hello")
	}
}
