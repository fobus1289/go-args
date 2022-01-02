```go
//example 
package main

import (
	go_args "github.com/fobus1289/go-args"
	"log"
)

var args = map[string]string{}

var keysArgs = []string{"port", "address"}

func main() {
	//ignore if not
	getArgs(false)
	//check all or ret error
	getArgs(true)

}

func getArgs(setError bool) {
	var err error
	args, err = go_args.GetMapArgs(setError, keysArgs...)

	if err != nil {
		panic(err)
	}

	log.Println(args)
}

```