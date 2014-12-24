// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"text/tabwriter"

	"github.com/millere/todo"
)

var cmdList = &Command{
	UsageLine: "list",
	Short:     "lists todos",
	Run:       runList,
}

func runList(cmd *Command, conf config, args []string) {
	todoFile, err := os.Open(conf.Todos)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			fmt.Println("No tasks :)")
			fmt.Print(fortune())
		default:
			fmt.Println(err)
		}
		return
	}
	todos, err := todo.FromReader(todoFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	//for _, t := range todos {
	//fmt.Printf("%s\n\n%#v\n\n", t.Raw, t)
	//}
	listPretty(todos)
}

func fortune() string {
	cmd := exec.Command("fortune")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}

func listPretty(ts todo.TaskList) {
	tw := tabwriter.NewWriter(
		os.Stdout,
		2,   // minwidth
		2,   // tabwidth
		4,   // padding
		' ', // tabchar
		0,   // flags
	)
	sort.Sort(ts)
	fmt.Fprintln(tw, "done\ttitle\tdue\tstart\t")
	for i := range ts {
		fmt.Fprintln(tw, ts[i])
	}
	tw.Flush()

}
