// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package todo implements a parser for my minimal todo list format.
package todo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

// A TaskList is a list of tasks
type TaskList []Task

func (l TaskList) Len() int      { return len(l) }
func (l TaskList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l TaskList) Less(i, j int) bool {
	// sort by:
	// not done before done
	// then by due date
	// then by start date
	// then alphabetically
	if l[i].Done && !l[j].Done {
		return false
	}
	if !l[i].Done && l[j].Done {
		return true
	}

	dbefore, eq := before(l[i].Due, l[j].Due)
	if !eq {
		return dbefore
	}

	sbefore, eq := before(l[i].Start, l[j].Start)
	if !eq {
		return sbefore
	}

	return l[i].Title < l[j].Title
}

// returns before, equal
func before(a, b time.Time) (bool, bool) {
	if a.Equal(b) {
		return false, true
	}
	if a.IsZero() && !b.IsZero() {
		return false, false
	}
	if !a.IsZero() && b.IsZero() {
		return true, false
	}
	if a.Before(b) {
		return true, false
	} else {
		return false, false
	}
}

func FromReader(r io.Reader) (TaskList, error) {
	s := bufio.NewScanner(r)
	var ret TaskList
	lno := 1
	for s.Scan() {
		line := s.Text()
		todo, err := Parse(line)
		if err != nil {
			return nil, fmt.Errorf("%v on line %v", err, lno)
		}
		todo.index = lno
		ret = append(ret, todo)
		lno++
	}
	return ret, nil
}

// A Task is represents a item in a todo list
type Task struct {
	Title    string
	Start    time.Time
	Due      time.Time
	Tags     []string
	Contexts []string
	index    int // line in file
	Raw      string
	Done     bool

	original string
}

// DateFormat is YY-MM-DD, with no times, time zone, etc.
const DateFormat = "2006-1-2"

// Parse takes a string and parses it as todo.txt formatted todo item
func Parse(r string) (Task, error) {
	if len(r) == 0 {
		return Task{}, errors.New("todo: parse empty string")
	}

	t := Task{Raw: r}
	tokens := strings.Fields(r)
	if len(tokens) == 0 {
		return Task{}, errors.New("todo: parse only whitespace")
	}

	if len(tokens) > 0 && tokens[0] == "x" {
		t.Done = true
		tokens = tokens[1:]
	}

	if len(tokens) == 0 {
		return Task{}, errors.New("todo: line contains only completion marker")
	}

	for _, token := range tokens {
		date, err := time.ParseInLocation(DateFormat, token, time.Local)
		switch {
		case err == nil:
			t.Due = date
		case strings.HasPrefix(token, "@"):
			if len(token[1:]) > 0 {
				t.Contexts = append(t.Contexts, token[1:])
			} else {
				t.Title = addToTitle(t.Title, token)
			}
		case strings.HasPrefix(token, "+"):
			if len(token[1:]) > 0 {
				t.Tags = append(t.Tags, token[1:])
			} else {
				t.Title = addToTitle(t.Title, token)
			}
		case strings.HasPrefix(token, "s:"):
			start, err := time.ParseInLocation(DateFormat, token[2:], time.Local)
			if err == nil {
				t.Start = start
			} else {
				t.Title = addToTitle(t.Title, token)
			}
		default:
			t.Title = addToTitle(t.Title, token)
		}
	}
	return t, nil
}

func addToTitle(title string, a string) string {
	if len(title) > 0 {
		title += " "
	}
	return title + a
}

func (t Task) String() string {
	done := ""
	if t.Done {
		done = "x"
	}
	due := ""
	if !t.Due.IsZero() {
		due = t.Due.Format(DateFormat)
	}
	start := ""
	if !t.Start.IsZero() {
		start = t.Start.Format(DateFormat)
	}
	out := fmt.Sprintf(
		"%v\t%v\t%v\t%v\t%v\t",
		t.index,
		done,
		t.Title,
		due,
		start,
	)

	return out
}
