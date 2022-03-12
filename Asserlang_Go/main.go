package main

import (
	"bufio"
	"fmt"
	"os"

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
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		tokenizer := token.NewTokenizer(line, token.TokenTypeDict)
		iter := 0
		for tokenizer.Next() == nil {
			iter++
		}
		fmt.Println(line)
		fmt.Println(tokenizer.Tokens)
	}

}
