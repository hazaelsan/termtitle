// Package formatter defines a terminal title formatter.
package formatter

import (
	"errors"
	"fmt"
	"sync"
)

// Formatter is a terminal title formatter.
type Formatter interface {
	// Start returns a control sequence required to start setting the terminal title.
	Start() string
	// Stop returns a control sequence required to stop setting the terminal title.
	Stop() string
	// Format returns a string formatted with start/stop control sequences.
	Format(string) string
}

var (
	// ErrNotFound is returned when the requested Formatter was not found in the registry.
	ErrNotFound = errors.New("formatter not found")
)

// Func is a function that returns a Formatter.
type Func func() Formatter

var registry = make(map[string]Func)
var mu sync.RWMutex

// Register adds a Formatter to the registry.  If Register is called twice with the same name, it panics.
func Register(name string, f Func) {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := registry[name]; ok {
		panic(fmt.Sprintf("formatter: Register called twice for %v", name))
	}
	registry[name] = f
}

// New returns a new Formatter with the given name.
func New(name string) (Formatter, error) {
	f, ok := registry[name]
	if !ok {
		return nil, ErrNotFound
	}
	return f(), nil
}

// List returns a list of Formatter names in the registry.
func List() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, len(registry))
	i := 0
	for n := range registry {
		names[i] = n
		i++
	}
	return names
}

// Format returns a string formatted with start/stop control sequences.
func Format(f Formatter, s string) string {
	return fmt.Sprintf("%s%s%s", f.Start(), s, f.Stop())
}
