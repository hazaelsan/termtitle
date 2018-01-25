// Package xterm defines a Formatter for XTerm variants.
package xterm

import "github.com/hazaelsan/termtitle/formatter"

func init() {
	formatter.Register("xterm", func() formatter.Formatter { return XTerm{} })
}

// XTerm is a Formatter for XTerm variants.
type XTerm struct{}

// Start returns a control sequence required to start setting the terminal title.
func (x XTerm) Start() string {
	return "\033]0;"
}

// Stop returns a control sequence required to stop setting the terminal title.
func (x XTerm) Stop() string {
	return "\007"
}

// Format returns a string formatted with start/stop control sequences.
func (x XTerm) Format(s string) string {
	return formatter.Format(x, s)
}
