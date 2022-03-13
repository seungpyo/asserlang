package token

import (
	"encoding/json"
	"fmt"
)

type TokenType uint

const (
	NONE = iota + 1
	EOL
	INCR
	DECR
	MUL
	ID
	SPACE
	START_PROGRAM
	END_PROGRAM
	DECL_INT
	ASSIGN_INT
	DECL_ASCII
	ASSIGN_ASCII
	STDIN
	STDOUT
	DECL_FUNCTION
	CALL_FUNCTION
	RETURN
	CONDITION
	CONDITION_EXEC
	GOTO
)

var TokenTypeDict = map[string]TokenType{
	"ㅋ":     INCR,
	"ㅎ":     DECR,
	"ㅌ":     MUL,
	"~":     SPACE,
	"쿠쿠루삥뽕": START_PROGRAM,
	"슉슈슉슉":  END_PROGRAM,
	"어쩔":    DECL_INT,
	"저쩔":    ASSIGN_INT,
	"우짤래미":  DECL_ASCII,
	"저짤래미":  DECL_INT,
	"ㅌㅂ":    STDIN,
	"ㅇㅉ":    STDOUT,
	"안물":    DECL_FUNCTION,
	"안궁":    CALL_FUNCTION,
	"무지개반사": RETURN,
	"화났쥬?":  CONDITION,
	"킹받쥬?":  CONDITION_EXEC,
	";;":    GOTO,
}

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
