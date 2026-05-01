package repl

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplStart(t *testing.T) {
	type testCase struct {
		name           string
		input          []string
		expectedOutput []string
		expectedErrs   []string
	}

	testCases := []testCase{
		{
			name: "test REPL can evaluate simple expressions",
			input: []string{
				"var x = 10;",
				"x;",
			},
			expectedOutput: []string{
				"10",
				"10",
			},
			expectedErrs: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inReader, inWriter := io.Pipe()
			outReader, outWriter := io.Pipe()

			r := New(inReader, outWriter, nil)
			go r.Start()

			// drain the ASCII art + welcome message + first prompt
			drainUntilPrompt(t, outReader)

			for i, cmd := range tc.input {
				// Send the command
				_, err := fmt.Fprintf(inWriter, "%s\n", cmd)
				assert.NoError(t, err)

				// read output until next prompt
				output := readUntilPrompt(t, outReader)
				output = strings.TrimSpace(output)

				assert.Equal(t, tc.expectedOutput[i], output)
			}

			inWriter.Close()
		})
	}
}

func readUntilPrompt(t *testing.T, r io.Reader) string {
	t.Helper()
	var buf strings.Builder
	b := make([]byte, 1)
	for {
		_, err := r.Read(b)
		if err != nil {
			return buf.String()
		}

		buf.WriteByte(b[0])
		if strings.HasSuffix(buf.String(), Prompt) {
			// strip the trailing prompt itself from the output
			s := buf.String()
			return s[:len(s)-len(Prompt)]
		}
	}
}

func drainUntilPrompt(t *testing.T, r io.Reader) {
	t.Helper()
	readUntilPrompt(t, r) // just discard
}
