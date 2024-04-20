package main

import (
	"log"
	"os"
	"strings"
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

	scriptString := string(scriptBytes)
	scriptString = strings.ReplaceAll("\r\n", "\n", scriptString)
	scriptString = strings.ReplaceAll("\r", "\n", scriptString)

	tokens := lexer.NewLexer(string(scriptBytes)).Parse()

	lexer.PrintTokens(tokens)
}
