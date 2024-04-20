package main

import (
	"fmt"
	"log"
	"os"
	"tlang/src/lexer"
)

func main() {
	scriptPath := ""

	if len(os.Args) < 2 {
		log.Fatalln("Provide a script file to execute")
	}

	scriptPath = os.Args[1]

	scriptBytes, err := os.ReadFile(scriptPath)
	if err != nil {
		log.Fatalln(err)
	}

	tokens := lexer.NewLexer(string(scriptBytes)).Parse()

	fmt.Printf("%#v\n", tokens)
}
