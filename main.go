package main

import (
	"fmt"
)

func main() {

	/*
		input, err := os.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
	*/

	input := " ( + 13 2  )"
	tokens := Lex(string(input))

	fmt.Print("Tokens : ")
	for _, token := range tokens {
		fmt.Print(token)
		fmt.Print(" ")
	}
	fmt.Println()

	fmt.Print("AST : ")
	ast, _ := Parse(tokens, 0)

	fmt.Println(ast.pretty())

}
