// Package null defines a null formatter, it does not output anything.
package null

import "github.com/hazaelsan/termtitle/formatter"

func init() {
	formatter.Register("null", func() formatter.Formatter { return Null{} })
}

// Null is a plain formatter, it does not output any control sequences.
type Null struct{}

// Start returns a control sequence required to start setting the terminal title.
func (n Null) Start() string {
	return ""
}

// Stop returns a control sequence required to stop setting the terminal title.
func (n Null) Stop() string {
	return ""
}

// Format returns a string formatted with start/stop control sequences.
func (n Null) Format(string) string {
	return ""
}
