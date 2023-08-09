package main

import (
	"fmt"

	"example.go/src"
)

func main() {

	//	input, err := os.ReadFile(os.Args[1])
	//	if err != nil {
	//		panic(err)
	//	}

	input := "(+ 1 3)"
	tokens := src.Lex(string(input))
	// fmt.Println(tokens[0])

	for i := 0; i < len(tokens); i++ {
		fmt.Println(tokens[i])
	}

}
