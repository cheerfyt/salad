package swgo

import (
	"fmt"
	"math"
	"os"
	"regexp"
	dbg "runtime/debug"
	"strconv"
	"strings"
)

var colors = []int{6, 2, 3, 4, 5, 1}
var color0 = "\u001b[0m"
var color3 = "\u001b[3"
var stack []byte

var names []string
var skips []string
var instance []*debugger

//debugger
type debugger struct {
	namespace   string
	enabled     bool
	color       int
	useColor    bool
	colorPrefix string
}

func (dg *debugger) setName(name string) *debugger {
	dg.namespace = name
	return dg
}

func (dg *debugger) setColor(color int) *debugger {
	dg.color = color
	return dg
}

func (dg *debugger) selectColor() *debugger {
	hash := 0
	for index := range dg.namespace {
		hash = ((hash << 5) - hash) + int(charCodeAt(dg.namespace, index))
		hash |= 0
	}
	dg.color = int(math.Abs(float64(hash))) % len(colors)
	return dg
}

func (dg *debugger) setColorPrefix() *debugger {
	var colorPrefix string
	if dg.color < 8 {
		colorPrefix = color3 + strconv.Itoa(dg.color)
	} else {
		colorPrefix = color3 + "8;5;" + strconv.Itoa(dg.color)
	}
	tmp := indent(2) + colorPrefix + ";1m" + dg.namespace + color0 + indent(1)
	dg.colorPrefix = tmp
	return dg
}

////////////////////////////////////////////////////////////////////////////////
func indent(num int) string {
	return strings.Repeat(" ", num)
}

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

func ignoreCase(re string) string {
	if !strings.HasPrefix(re, "(?i)") {
		re = "(?i)" + re
	}
	return re
}

func save(name string) error {
	return os.Setenv("SWGO", name)
}

func load() string {
	return os.Getenv("SWGO")
}

func enable(namespace string) {
	save(namespace)
	re1 := regexp.MustCompile(`[\s,]+`)
	re2 := regexp.MustCompile(ignoreCase("\\*"))
	split := re1.FindAllString(namespace, -1)
	for _, value := range split {
		fmt.Println(value)
		if value == "" {
			continue
		}
		str := re2.ReplaceAllString(value, ".*?")
		if string(str[0]) == "-" {
			skips = append(skips, "^"+str[1:]+"$")
		} else {
			names = append(names, "^"+str+"$")
		}
	}
}

func enabled(namespace string) bool {
	if namespace == "" {
		return false
	}
	if string(namespace[len(namespace)-1]) == "*" {
		return true
	}
	for _, value := range names {
		if ok, _ := regexp.MatchString(value, namespace); ok {
			return true
		}
	}
	for _, value := range skips {
		if ok, _ := regexp.MatchString(value, namespace); ok {
			return false
		}
	}
	return false
}

func makeNewLine(format string) string {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	return format
}

//NewDebugger new instance
func NewDebugger(namespace string) func(format string, args ...interface{}) (int, error) {
	var debug = &debugger{}
	debug = debug.setName(namespace).selectColor().setColorPrefix()
	return func(format string, args ...interface{}) (int, error) {
		debug.enabled = enabled(load())
		if !debug.enabled {
			return 0, nil
		}
		dbg.PrintStack()
		format = makeNewLine(debug.colorPrefix + format)
		instance = append(instance, debug)
		return fmt.Printf(format, args...)
	}
}
