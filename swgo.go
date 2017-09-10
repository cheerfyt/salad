package swgo

import (
	"fmt"
	"math"
	"os"
	"path"
	"regexp"
	dbg "runtime/debug"
	"strconv"
	"strings"
)

var colors = []int{6, 2, 3, 4, 5, 1}

var colorsMore = []int{
	31, 32, 33, 34, 35, 36, 37,
	90, 91, 92, 93, 94, 95, 96, 97,
}

var color0 = "\u001b[0m"
var color3 = "\u001b[3"

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
	dg.color = colorsMore[int(math.Abs(float64(hash)))%len(colors)]
	return dg
}

func (dg *debugger) setColorPrefix(stackInfo string) *debugger {
	var colorPrefix string
	colorPrefix = color3 + "8;5;" + strconv.Itoa(dg.color)
	dg.colorPrefix = "  " + colorPrefix + ";1m" + dg.namespace + stackInfo + color0 + " "
	return dg
}

////////////////////////////////////////////////////////////////////////////////
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
	split := re1.Split(namespace, -1)
	for _, value := range split {
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

	for _, value := range instance {
		value.enabled = enabled(value.namespace)
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

// =============================================================================
func getPathInfo(project string) string {
	_, file := path.Split(project)
	if file != "" {
		return strings.Split(file, " ")[0]
	}
	return file
}

func getFuncName(long string) string {
	return strings.Split(long, ".")[1]
}

func buildRunInfo(where string, fname string) string {
	return fmt.Sprintf(" [info:%s-%s]", where, fname)
}

// =============================================================================

//NewDebugger new instance
func NewDebugger(namespace string) func(format string, args ...interface{}) (int, error) {
	var debug = &debugger{}
	debug = debug.setName(namespace).selectColor()
	instance = append(instance, debug)
	enable(load())
	return func(format string, args ...interface{}) (int, error) {
		if !debug.enabled {
			return 0, nil
		}
		/* Get call stack info */
		split := strings.Split(string(dbg.Stack()), "\n")
		stackInfo := buildRunInfo(getPathInfo(split[6]), getFuncName(split[5]))
		debug = debug.setColorPrefix(stackInfo)
		format = makeNewLine(debug.colorPrefix + format)
		return fmt.Printf(format, args...)
	}
}
