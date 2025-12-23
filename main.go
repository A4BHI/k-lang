package main

import (
	"fmt"
	klang "k/src"
)

func main() {
	l := klang.NewLexer("make x = 10;")
	p := klang.NewParser(l)
	pg := p.ParseProgram()
	fmt.Println("Errors:", p.Errors)
	fmt.Println("Program:", pg)
	ms := pg.Statements[0].(*klang.MakeStatement)

	fmt.Println(ms.Name.Value)
	fmt.Println(ms.Value.(*klang.IntegerLiteral).Value)

	for i, stmt := range pg.Statements {
		fmt.Printf("Statement: %d: %#v\n", i, stmt)
	}

}
