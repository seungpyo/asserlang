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
		tokenizer := token.NewTokenizer(line)
		tokenizer = *tokenizer.
			Reserve("ㅋ", token.INCR).
			Reserve("ㅎ", token.DECR).
			Reserve("ㅌ", token.MUL).
			Reserve("쿠쿠루삥뽕", token.START_PROGRAM).
			Reserve("슉슈슉슉", token.END_PROGRAM)
		iter := 0
		for tokenizer.Next() == nil {
			iter++
		}
		fmt.Println(line, tokenizer.Tokens)
	}

}
