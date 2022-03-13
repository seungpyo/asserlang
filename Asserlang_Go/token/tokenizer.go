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
	Identifiers map[string]bool
	line        []rune
	pos         int
	length      int
}

func NewTokenizer(reserved map[string]TokenType) Tokenizer {
	return Tokenizer{
		TokenBuffer: "",
		Tokens:      []Token{},
		Reserved:    reserved,
		Identifiers: map[string]bool{},
		line:        []rune(""),
		pos:         0,
		length:      0,
	}
}

func (t *Tokenizer) String() string {
	s, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Marshalizing Tokenizer:", err)
	}
	return string(s)
}

func (t *Tokenizer) Feed(line string) {
	t.line = []rune(line)
	t.pos = 0
	t.length = utf8.RuneCountInString(line)
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
		t.Append(NewToken("", EOL))
		return errors.New("end of line")
	}

	t.TokenBuffer += string(t.line[t.pos])

	// Case 1. Current buffer is a reserved word
	bufferType := t.TypeQuery(t.TokenBuffer)
	if bufferType != ID {
		t.Append(NewToken(t.TokenBuffer, bufferType))
		t.TokenBuffer = ""
		t.pos++
		return nil
	}

	// Case 2. Current buffer is an identifier of a symbol
	_, exists := t.Identifiers[t.TokenBuffer]
	if exists {
		t.Append(NewToken(t.TokenBuffer, ID))
		t.TokenBuffer = ""
		t.pos++
		return nil
	}

	// Case 3. End of line (e.g. ㅇㅉ냉장고)
	if t.pos+1 >= t.length {
		t.Identifiers[t.TokenBuffer] = true
		t.Append(NewToken(t.TokenBuffer, ID))
		t.TokenBuffer = ""
		t.pos++
		return nil
	}

	// Case 4. The next character is a 1-byte keyword
	// In this case, flush current buffer as an identifier
	nextChar := string(t.line[t.pos+1])
	peakedType := t.TypeQuery(nextChar)
	if peakedType != ID {
		t.Identifiers[t.TokenBuffer] = true
		t.Append(NewToken(t.TokenBuffer, ID))
		t.TokenBuffer = ""
		t.pos++
		return nil
	}

	t.pos++
	return nil
}
