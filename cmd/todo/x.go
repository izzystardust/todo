// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/millere/todo"
)

var cmdX = &Command{
	UsageLine: "x",
	Short:     "x marks a task as completed",
	Run:       runX,
}

func runX(cmd *Command, conf config, args []string) {
	var toMark []int
	for _, a := range args {
		mark, err := strconv.Atoi(a)
		if err != nil {
			fmt.Println("arguments to x must be integers")
			return
		}
		toMark = append(toMark, mark)
	}
	todoFile, _ := os.Open(conf.Todos)
	todos, err := todo.FromReader(todoFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	didWork := false
	for i := range todos {
		//BUG(@millere): off-by-one error
		if elementOf(toMark, i+1) {
			todos[i].Done = true
			didWork = true
		}
	}

	if didWork {
		file, err := os.Create(conf.Todos)
		if err != nil {
			fmt.Println("Couldn't modify todo file:", err)
			return
		}
		// TODO: if -s sort before writing
		for i := range todos {
			fmt.Fprintln(file, todos[i].UnParse())
		}
	}
}

func elementOf(as []int, a int) bool {
	for _, b := range as {
		if b == a {
			return true
		}
	}
	return false
}
