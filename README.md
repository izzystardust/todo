# Todo

Todo is a todo list manager that works with a simple text format.

Ideally, this is usable without the command. The command just helps.

## The Format

### Why not todo.txt?

It has more than I need or will use.

### Examples:

```
Take out trash
2014-12-23 Feed cats
Write novel 2015-12-31 s:2015-12-30
x Eat lunch
```

### The Specification

A todo is a single line of utf8 text. All whitespace characters are treated as spaces.
Each whitespace separated string is treated as a token. The rules for parsing are as follows:

- If the first token of the file is an `x` lower case x, the task is completed.
- If the token matches the date format `YYYY-MM-DD`, it is the due date of the task.
- If the token matches the date format `s:YYYY-MM-DD`, it is the schedules start date of the task.
- If the token starts with `+` and `len(token) > 1`, the token specifies a case-insensitive tag.
- If the token starts with `@` and `len(token) > 1`, the token specifies a case-insensitive context.
- Otherwise, the token is part of the title of the task.

## The Command

### Install

To install from source, you'll need [Go](https://golang.org) installed. With a properly configured
`GOPATH`, run `go get github.com/millere/todo/...`. The compiled `todo` binary will be placed in `$GOPATH/bin`.

### Usage

NOTE: THIS IS IN EARLY DEVELOPMENT. Stuff probably doesn't work right.

## The Library

[![GoDoc](https://godoc.org/github.com/millere/todo?status.svg)](https://godoc.org/github.com/millere/todo)

As much functionality of the command as possible has been implemented in the library.

