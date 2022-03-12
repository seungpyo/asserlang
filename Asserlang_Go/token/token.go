package token

import (
	"encoding/json"
	"fmt"
)

type TokenType uint

const (
	NONE = iota + 1
	INCR
	DECR
	MUL
	ID
	DELIMETER
	START_PROGRAM
	END_PROGRAM
)

type Token struct {
	Raw  string
	Type TokenType
	Data interface{}
}

func NewToken(raw string, tokenType TokenType) *Token {
	t := Token{
		Raw:  raw,
		Type: tokenType,
		Data: nil,
	}
	return &t
}

func (t *Token) String() string {
	s, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Marshalizing token: ", err)
	}
	return string(s)
}
