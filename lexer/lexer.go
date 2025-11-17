package main

import "fmt"

const (
	MAKE = "MAKE"
)

var Tokens = map[string]string{
	"make": MAKE,
}

func OutPut(src string) {
	token, ok := Tokens[src]

	if ok {
		fmt.Print(token)
	} else {
		fmt.Print("IDENT")
	}

}

func main() {
	// OutPut("make")
	OutPut("x")
}
