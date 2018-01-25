// Package screen defines a Formatter for GNU Screen, compatible with tmux.
package screen

import "github.com/hazaelsan/termtitle/formatter"

func init() {
	formatter.Register("screen", func() formatter.Formatter { return Screen{} })
}

// Screen is a Formatter for GNU Screen.
type Screen struct{}

// Start returns a control sequence required to start setting the terminal title.
func (s Screen) Start() string {
	return "\033k"
}

// Stop returns a control sequence required to stop setting the terminal title.
func (s Screen) Stop() string {
	return "\033\\"
}

// Format returns a string formatted with start/stop control sequences.
func (s Screen) Format(ss string) string {
	return formatter.Format(s, ss)
}
