// Package auto is a pseudo-Formatter returning a suitable Formatter based on the underlying terminal.
package auto

import (
	"os"
	"strings"

	"github.com/hazaelsan/termtitle/formatter"
	"github.com/hazaelsan/termtitle/formatter/screen"
	"github.com/hazaelsan/termtitle/formatter/xterm"
)

func init() {
	if f, err := New(); err == nil {
		formatter.Register("auto", func() formatter.Formatter { return f })
	}
}

// New returns a suitable Formatter based on the underlying terminal's characteristics.
func New() (formatter.Formatter, error) {
	term := os.Getenv("TERM")
	if strings.HasPrefix(term, "screen") {
		return screen.Screen{}, nil
	}
	// TODO: Check for cases when the null formatter should be returned.
	return xterm.XTerm{}, nil
}
