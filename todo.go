// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package todo implements to do items in a GTD-compatible style.
// It includes the ability to parse the todo.txt format.
package todo

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// A Todo is represents a item in a todo list
type Todo struct {
	Title     string
	Created   time.Time
	Start     time.Time
	Due       time.Time
	Completed time.Time
	Tags      []string
	Contexts  []string
	Pri       Priority
	Done      bool

	original string
}

// A Priority is the priority of a todo item
type Priority int

const (
	// NoPriority represents the zero value of a priority
	// A priority of zero is the same as not having a proirity
	NoPriority Priority = iota
)

// ByPriority implements sort.Interface to sort Todos by priority
type ByPriority []Todo

func (a ByPriority) Len() int      { return len(a) }
func (a ByPriority) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool {
	if a[i].Pri == NoPriority {
		return false
	}
	if a[j].Pri == NoPriority {
		return true
	}
	return a[i].Pri < a[j].Pri
}

// DateFormat is YYYY-MM-DD, with no times, time zone, etc.
const DateFormat = "2006-1-2"

type leader uint8

// The *Leader constants are used to classify todo tags into the approrpiate
// category
const (
	ContextLeader = '@'
	ProjectLeader = '+'
)

func addTag(token string, tagList *[]string) error {
	if len(token[1:]) == 0 {
		return fmt.Errorf("todo: contains empty context")
	}
	*tagList = append(*tagList, token[1:])
	return nil
}

// Parse takes a string and parses it as todo.txt formatted todo item
func Parse(r string) (Todo, error) {
	tokens := strings.Fields(r)
	if len(tokens) == 0 {
		return Todo{}, fmt.Errorf("todo: parse empty string")
	}
	var t Todo
	var token string
	token, tokens = pop(tokens)

	// Parse completion rules:
	// 1. Completed tasks start with an x and a space
	// 2. The x may be followed by a completion timestamp
	if token == "x" {
		t.Done = true

		if len(tokens) < 1 {
			return Todo{}, fmt.Errorf("todo: contains only done marker")
		}
		token, tokens = pop(tokens)
		completed, err := time.ParseInLocation(DateFormat, token, time.Local)
		if err == nil {
			// parsed time successfully, we have a completion stamp
			t.Completed = completed

			fmt.Println(r, len(tokens))
			if len(tokens) < 1 {
				return Todo{}, fmt.Errorf("todo: contains only done marker and completion time")
			}
			token, tokens = pop(tokens)
		}
	}

	// Parse priority rules
	// 1. A priority, if present, must be first after completion marks
	// 2. A priority is of the form "(x)", where x is a letter that is the priority
	// 3. A priority may be followed by a completion date
	isPriority, err := regexp.MatchString(`[(][[:alpha:]][)]$`, token)
	if err != nil {
		panic(err)
	}
	if isPriority {
		t.Pri = Priority(unicode.ToLower(rune(tokens[0][1]))) - 'a' + 1
	}

	for _, token := range tokens {
		switch {
		case token[0] == ContextLeader:
			addTag(token, &t.Contexts)
			if err != nil {
				return Todo{}, err
			}
		case token[0] == ProjectLeader:
			addTag(token, &t.Tags)
			if err != nil {
				return Todo{}, err
			}
		default:
			if len(t.Title) > 0 {
				t.Title += " "
			}
			t.Title += token
		}

	}
	return t, nil
}

func pop(as []string) (string, []string) {
	return as[0], as[1:]
}

func (t Todo) String() string {
	now := time.Now()
	switch {
	case t.Start.Before(now) && t.Due.IsZero():
		return t.Title
	case t.Start.Before(now):
		fmt.Sprintf("%v (due %v)", t.Title, t.Due)
	default:
		fmt.Sprintf("%v (postponed until %v, due %v)", t.Title, t.Start, t.Due)
	}
	panic("Wat")
}
