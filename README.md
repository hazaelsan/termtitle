# termtitle

[![GoDoc](https://godoc.org/github.com/hazaelsan/termtitle?status.svg)](https://godoc.org/github.com/hazaelsan/termtitle)

Sets the terminal title.

This can be used as a helper for various CLI-based applications (e.g., `mutt`)
to set a context-aware terminal title.

## Installation

```shell
go get -u github.com/hazaelsan/termtitle
```

## Usage

```shell
$GOBIN/termtile -title_auto "my fancy title"
```

## Formatters

There are several formatters for setting the terminal title.

### auto

Pseudo-formatter, if a `screen`/`tmux` terminal is detected it returns the
`screen` formatter, otherwise it returns a `xterm` formatter.

### null

The bitbucket, doesn't output anything.

### raw

A formatter that doesn't output any control sequences.

### screen

Uses `screen`/`tmux` compatible control sequences.

### xterm

Uses `XTerm` compatible control sequences.
