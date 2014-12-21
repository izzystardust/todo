// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/millere/todo"
)

var cmdList = &Command{
	UsageLine: "list",
	Short:     "lists todos",
	Run:       runList,
}

func runList(cmd *Command, args []string) {
	todoFile, err := os.Open("todo.txt")
	if err != nil {
		panic(err)
	}
	s := bufio.NewScanner(todoFile)
	for s.Scan() {
		line := s.Text()
		fmt.Println(line)
		todo, err := todo.Parse(line)
		if err != nil {
			println(err)
		}
		fmt.Printf("%#v\n\n", todo)
	}
	fmt.Println("I would list my todos now")
}
