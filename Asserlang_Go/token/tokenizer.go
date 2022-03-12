package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"unicode/utf8"
)

type Tokenizer struct {
	TokenBuffer string
	Tokens      []Token
	Reserved    map[string]TokenType
	line        []rune
	pos         int
	length      int
}

func NewTokenizer(line string, reserved map[string]TokenType) Tokenizer {
	return Tokenizer{
		TokenBuffer: "",
		Tokens:      []Token{},
		Reserved:    reserved,
		line:        []rune(line),
		pos:         0,
		length:      utf8.RuneCountInString(line),
	}
}

func (t *Tokenizer) String() string {
	s, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Marshalizing Tokenizer:", err)
	}
	return string(s)
}

func (t *Tokenizer) Append(tok *Token) {
	t.Tokens = append(t.Tokens, *tok)
}

func (t *Tokenizer) TypeQuery(key string) TokenType {
	tokenType, exists := t.Reserved[key]
	if exists {
		return tokenType
	} else {
		return ID
	}
}

func (t *Tokenizer) Next() error {
	if t.pos >= t.length {
		return errors.New("end of line")
	}

	// Case 1. Current buffer is a reserved word
	t.TokenBuffer += string(t.line[t.pos])
	bufferType := t.TypeQuery(t.TokenBuffer)
	if bufferType != ID {
		t.Append(NewToken(t.TokenBuffer, bufferType))
		t.TokenBuffer = ""
		t.pos++
		return nil
	}

	// Case 2. End of line (e.g. ㅇㅉ냉장고)
	if t.pos+1 >= t.length {
		t.Append(NewToken(t.TokenBuffer, ID))
		t.pos++
		return nil
	}
	nextChar := string(t.line[t.pos+1])
	peakedType := t.TypeQuery(nextChar)
	if peakedType != ID {
		t.Append(NewToken(t.TokenBuffer, ID))
		t.TokenBuffer = ""
		t.pos++
		return nil
	}

	t.pos++
	return nil
}
