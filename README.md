# salad
golang debugger like npm debug package

## usage

- install: `go get github.com/cheeryt/salad`
- test: `make test`

```go
package main

import "github.com/cheeryt/salad"

func main() {
  debug1 := salad.NewDebugger("salad:app")
  debug2 := salad.NewDebugger("salad:test")
  for i:=0; i < 2; i ++ {
    debug1("hello, salad app")
    debug2("hello, salad test")
  }
}
```

run go code:

```sh
$ SALAD=salad:* go run main.go
$ SALAD=salad:app go run main.go
$ SALAD=salad:test go run main.go
```
