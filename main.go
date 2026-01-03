package main

import (
	"fmt"
	klang "k/src"
)

func main() {
	l := klang.NewLexer("foo(1,2)")
	p := klang.NewParser(l)
	pg := p.ParseProgram()
	fmt.Println("Errors:", p.Errors)
	fmt.Println("Program:", pg)
	ms := pg.Statements[0].(*klang.ExpressionStatement)

	// fmt.Println(ms.Name.Value)
	fmt.Println(ms.Expression.(*klang.FunctionCall).Function)
	fmt.Println(ms.Expression.(*klang.FunctionCall).Arguments[0].(*klang.IntegerLiteral).Value)
	fmt.Println(ms.Expression.(*klang.FunctionCall).Arguments[1].(*klang.IntegerLiteral).Value)

	for i, stmt := range pg.Statements {
		fmt.Printf("Statement: %d: %#v\n", i, stmt)
	}

}
