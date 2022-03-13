package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/seungpyo/asserlang/astv_parser"
	"github.com/seungpyo/asserlang/token"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing input")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Opening", os.Args[1], err)
		return
	}
	scanner := bufio.NewScanner(f)
	lineNum := 0
	tokenizer := token.NewTokenizer(token.TokenTypeDict)
	for scanner.Scan() {
		line := scanner.Text()
		tokenizer.Feed(line)
		lineNum++
		iter := 0
		for tokenizer.Next() == nil {
			iter++
		}
	}
	// fmt.Println(tokenizer.Tokens)

	fmt.Println(tokenizer.Tokens)
	parser := astv_parser.NewParser(tokenizer.Tokens)
	for parser.NextLine() == nil {
		fmt.Println(parser.CurrentLine())
	}
}
