package repl

import (
	"bufio"
	"io"

	"github.com/OJOMB/donkey/internal/lexer"
	"github.com/OJOMB/donkey/internal/tokens"
	"github.com/OJOMB/donkey/pkg/logs"
)

const Prompt = ">> "

type Repl struct {
	in  io.Reader
	out io.Writer

	logger logs.Logger
}

func New(in io.Reader, out io.Writer, l logs.Logger) *Repl {
	if l == nil {
		l = logs.NewNullLogger()
	}

	return &Repl{in: in, out: out, logger: l.With("component", "repl")}
}

func (r *Repl) Start() {
	// create a new scanner to read input from the user
	scanner := bufio.NewScanner(r.in)

	if _, err := r.out.Write([]byte("Welcome to the Donkey programming language!\n")); err != nil {
		return
	}

	for {
		if _, err := r.out.Write([]byte(Prompt)); err != nil {
			r.logger.Error("failed to write prompt", "error", err)
			return
		}

		scanned := scanner.Scan()
		if !scanned {
			if err := scanner.Err(); err != nil {
				r.logger.Error("failed to read input", "error", err)
			}

			return
		}

		line := scanner.Text()
		l := lexer.New(line, r.logger)

		for tok := l.NextToken(); tok.Type != tokens.TypeEOF; tok = l.NextToken() {
			if _, err := r.out.Write([]byte(tok.Lexeme + "\n")); err != nil {
				r.logger.Error("failed to write token", "error", err)

				return
			}
		}
	}
}
