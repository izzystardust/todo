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

Todo has four commands: add, list, x, and archive.

#### add

```
todo add update readme 2015-01-08 @computer
```

Add parses the given line as a todo to check syntax and adds it to the todo file.

#### list

```
todo list
todo list @computer
todo list readme
todo list -s
```

List lists todos. If words are given, it filters by matching all of those words. The `-s` flag sorts the results in a way I find useful. This `-a` flag causes completed tasks to be listed.

#### help

```
todo help
todo help list
```

Prints usage information.

#### x

```
todo x 3
```

X marks task n as completed.

### archive

```
todo archive
```

Archive moves all completed tasks from the todo file to the archive file.

## The Library

[![GoDoc](https://godoc.org/github.com/millere/todo?status.svg)](https://godoc.org/github.com/millere/todo)

As much functionality of the command as possible has been implemented in the library.

