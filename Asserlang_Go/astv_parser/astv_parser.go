package astv_parser

import (
	"errors"

	"github.com/seungpyo/asserlang/token"
)

type Parser struct {
	tokens    []token.Token
	lineStart int
	lineEnd   int
}

func NewParser(tokens []token.Token) Parser {
	return Parser{
		tokens:    tokens,
		lineStart: -1,
		lineEnd:   -1,
	}
}

func (p *Parser) NextLine() error {
	p.lineEnd++
	if p.lineEnd >= len(p.tokens) {
		return errors.New("end of tokens")
	}
	p.lineStart = p.lineEnd
	for ; p.tokens[p.lineEnd].Type != token.EOL; p.lineEnd++ {
	}
	return nil
}

func (p *Parser) CurrentLine() []token.Token {
	return p.tokens[p.lineStart:p.lineEnd]
}

type parser_fn func([]token.Token) error

func (p *Parser) ParseLine() {
	firstToken := p.tokens[p.lineStart]
	var parser_impl parser_fn
	switch firstToken.Type {
	case token.EOL:
		parser_impl = p.ParseEOL
	case token.INCR:
		parser_impl = p.ParseConstExpr
	case token.DECR:
		parser_impl = p.ParseConstExpr
	case token.MUL:
		parser_impl = p.ParseConstExpr
	case token.ID:
		parser_impl = p.ParseNYI
	case token.START_PROGRAM:
		parser_impl = p.ParseStartProgram
	case token.END_PROGRAM:
		parser_impl = p.ParseEndProgram
	case token.DECL_INT:
		parser_impl = p.ParseDeclareInt
	case token.ASSIGN_INT:
		parser_impl = p.ParseAssignInt
	case token.DECL_ASCII:
		parser_impl = p.ParseDeclareAscii
	case token.ASSIGN_ASCII:
		parser_impl = p.ParseAssignAscii
	case token.STDIN:
		parser_impl = p.ParseStdin
	case token.STDOUT:
		parser_impl = p.ParseStdout

	}
	parser_impl(p.tokens[p.lineStart:p.lineEnd])
}

func (p *Parser) ParseNYI(tokens []token.Token) error {
	return errors.New("nyi")
}

func (p *Parser) ParseEOL(tokens []token.Token) error {
	return nil
}

func (p *Parser) ParseConstExpr(tokens []token.Token) error {
	return nil
}

func (p *Parser) ParseID(tokens []token.Token) error {
	return nil
}

func (p *Parser) ParseStartProgram(tokens []token.Token) error {
	return nil
}

func (p *Parser) ParseEndProgram(tokens []token.Token) error {
	return nil
}

func (p *Parser) ParseDeclareInt(tokens []token.Token) error {
	return nil
}

func (p *Parser) ParseAssignInt(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseDeclareAscii(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseAssignAscii(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseStdin(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseStdout(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseDeclareFunction(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseCallFunction(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseReturn(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseCondition(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseConditionExecution(tokens []token.Token) error {
	return nil
}
func (p *Parser) ParseGoto(tokens []token.Token) error {
	return nil
}
