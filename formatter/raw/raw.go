// Package raw defines a plain formatter that does not output any control sequences.
package raw

import "github.com/hazaelsan/termtitle/formatter"

func init() {
	formatter.Register("raw", func() formatter.Formatter { return Raw{} })
}

// Raw is a plain formatter, it does not output any control sequences.
type Raw struct{}

// Start returns a control sequence required to start setting the terminal title.
func (r Raw) Start() string {
	return ""
}

// Stop returns a control sequence required to stop setting the terminal title.
func (r Raw) Stop() string {
	return ""
}

// Format returns a string formatted with start/stop control sequences.
func (r Raw) Format(s string) string {
	return s
}
