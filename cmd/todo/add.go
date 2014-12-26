// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/millere/todo"
)

var cmdAdd = &Command{
	UsageLine: "add",
	Short:     "adds a todo to the todo list",
	Run:       runAdd,
}

func runAdd(cmd *Command, conf config, args []string) {
	in := strings.Join(args, " ")
	_, err := todo.Parse(in)
	if err != nil {
		fmt.Println("Malformed task:", err)
		return
	}
	fi, err := os.OpenFile(conf.Todos, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0)
	if err != nil {
		fmt.Println("Wat:", err)
	}
	fmt.Fprintln(fi, in)
}
