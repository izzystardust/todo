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
}

func init() {
	cmdList.Run = runList
}

var listSorted = cmdList.Flag.Bool("s", false, "")
var listN = cmdList.Flag.Int("n", 0, "")

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

	for _, a := range args {
		todos = todos.Filter(a)
	}

	if *listSorted {
		sort.Sort(todos)
	}

	listPretty(todos, *listN)
}

func fortune() string {
	cmd := exec.Command("fortune")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}

func listPretty(ts todo.TaskList, max int) {
	tw := tabwriter.NewWriter(
		os.Stdout,
		2,   // minwidth
		2,   // tabwidth
		4,   // padding
		' ', // tabchar
		0,   // flags
	)
	if max == 0 {
		max = len(ts)
	}

	fmt.Fprintln(tw, "\tdone\ttitle\tdue\tstart\tcontexts\ttags\t")
	for i := range ts {
		if !ts[i].Done {
			fmt.Fprintln(tw, ts[i])
			if i >= max-1 {
				break
			}
		}
	}
	tw.Flush()

}
