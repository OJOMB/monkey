package evaluator

import (
	"fmt"

	"github.com/OJOMB/donkey/internal/objects"
)

func newError(format string, args ...any) *objects.ErrorValue {
	return &objects.ErrorValue{Message: fmt.Sprintf(format, args...)}
}
