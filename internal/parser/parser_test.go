package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/OJOMB/monkey/internal/ast"
	"github.com/OJOMB/monkey/internal/lexer"
	"github.com/OJOMB/monkey/internal/tokens"
)

func TestLetStmts(t *testing.T) {
	type testCase struct {
		Name           string
		input          string
		expectedOutput *ast.Program
	}

	var testCases = []testCase{
		{
			Name: "test let statements",
			input: `
				let x = 5;
				let y = 10;
				let __foobar__ = 838383;
			`,
			expectedOutput: &ast.Program{
				Statements: []ast.Statement{
					&ast.LetStatement{
						Token: tokens.Token{Type: "LET", Lexeme: "let"},
						Name: &ast.Identifier{
							Token: tokens.Token{Type: "IDENT", Lexeme: "x"},
							Value: "x",
						},
					},
					&ast.LetStatement{
						Token: tokens.Token{Type: "LET", Lexeme: "let"},
						Name: &ast.Identifier{
							Token: tokens.Token{Type: "IDENT", Lexeme: "y"},
							Value: "y",
						},
					},
					&ast.LetStatement{
						Token: tokens.Token{Type: "LET", Lexeme: "let"},
						Name: &ast.Identifier{
							Token: tokens.Token{Type: "IDENT", Lexeme: "__foobar__"},
							Value: "__foobar__",
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			p := New(lexer.New(tc.input))
			program := p.ParseProgram()
			require.NotNil(t, program)

			assert.Equal(t, tc.expectedOutput, program)
		})
	}
}
