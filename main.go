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

var (
	debug1 = swgo.NewDebugger("swgo:xiang")
	debug2 = swgo.NewDebugger("swgo:show")
	debug3 = swgo.NewDebugger("swgo:man")
	debug4 = swgo.NewDebugger("swgo:dir")
)

func test1() {
	debug3("test1")
}

func main() {
	for i := 0; i < 2; i++ {
		debug1("shangtian: %s", "hello")
	}
	debug2("swgo:show")
	debug4("swgo:dir")
	test1()
}
