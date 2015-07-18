package gl

import (
	"fmt"
	"runtime"
)

func IsSafe() bool {
	return safety
}

func CheckErr() {
	if err := GetError(); err != nil {
		var stack = []struct {
			file string
			line int
		}{}
		var file string
		var line int
		skip := 2
		ok := true
		for ok {
			_, file, line, ok = runtime.Caller(skip)
			skip++
			if ok {
				stack = append(stack, struct {
					file string
					line int
				}{file, line})
			}
		}
		stack = stack[0 : len(stack)-2] //remove the 2 go func call
		for x := len(stack) - 1; x >= 0; x-- {
			fmt.Printf("%s:%d\n", stack[x].file, stack[x].line)
		}
		fmt.Println(err)
	}
}
