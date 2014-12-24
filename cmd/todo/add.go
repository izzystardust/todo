// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"

var cmdAdd = &Command{
	UsageLine: "add",
	Short:     "adds a todo to the todo list",
	Run:       runAdd,
}

func runAdd(cmd *Command, conf config, args []string) {
	fmt.Println("Adding a todo. ")
}
