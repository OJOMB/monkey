package main

import (
	"os"

	"github.com/OJOMB/monkey/internal/repl"
)

func main() {
	r := repl.New(os.Stdin, os.Stdout)
	r.Start()
}
