// terminal_title is a utility that sets the terminal title.
// Outputting via /dev/tty is needed in order to make some applications work correctly (e.g., mutt).
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/hazaelsan/termtitle/formatter"
	_ "github.com/hazaelsan/termtitle/formatter/auto"
	_ "github.com/hazaelsan/termtitle/formatter/null"
	_ "github.com/hazaelsan/termtitle/formatter/raw"
	_ "github.com/hazaelsan/termtitle/formatter/screen"
	_ "github.com/hazaelsan/termtitle/formatter/xterm"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	force = flag.Bool("force", false, "force output even if output isn't connected to a terminal")
)

var (
	tty = "/dev/tty"
)

type fmtSet struct {
	name string
	f    formatter.Formatter
	flag *string
}

type writer struct {
	file *os.File
	fs   []*fmtSet
}

type writerSet struct {
	name   string
	suffix string
	writer writer
}

func makeWriters() ([]*writerSet, error) {
	writers := []*writerSet{
		&writerSet{
			name:   "stdout",
			writer: writer{file: os.Stdout},
		},
	}
	// /dev/tty may not be available in all platforms (e.g., Windows)
	if f, err := os.OpenFile(tty, os.O_WRONLY, os.ModeAppend); err == nil {
		w := &writerSet{
			name:   tty,
			suffix: "_tty",
			writer: writer{file: f},
		}
		writers = append(writers, w)
	}
	for _, w := range writers {
		for _, fName := range formatter.List() {
			v := ""
			f, err := formatter.New(fName)
			if err != nil {
				return nil, err
			}
			fs := &fmtSet{
				name: fName,
				f:    f,
				flag: &v,
			}
			name := fmt.Sprintf("title_%v%v", fName, w.suffix)
			usage := fmt.Sprintf("%v title via %v", fName, w.name)
			flag.StringVar(fs.flag, name, "", usage)
			w.writer.fs = append(w.writer.fs, fs)
		}
	}
	return writers, nil
}

func main() {
	writers, err := makeWriters()
	if err != nil {
		glog.Fatal(err)
	}
	flag.Parse()

	errors := 0
	for _, w := range writers {
		for _, fs := range w.writer.fs {
			if !*force && !terminal.IsTerminal(int(w.writer.file.Fd())) {
				glog.Errorf("not a terminal: %v", w.writer.file.Name())
				errors++
				continue
			}
			_, err := fmt.Fprintf(w.writer.file, fs.f.Format(*fs.flag))
			if err != nil {
				glog.Error(err)
				errors++
			}
		}
	}
	os.Exit(errors)
}
