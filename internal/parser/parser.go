package parser

import (
	"github.com/OJOMB/monkey/internal/ast"
	"github.com/OJOMB/monkey/internal/lexer"
	"github.com/OJOMB/monkey/internal/tokens"
)

type Parser struct {
	l *lexer.Lexer

	currToken tokens.Token
	peekToken tokens.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so currToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
