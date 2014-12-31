// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/millere/todo"
)

var cmdArchive = &Command{
	UsageLine: "archive",
	Short:     "archives completed todos",
}

func init() {
	cmdArchive.Run = runArchive
}

func runArchive(cmd *Command, conf config, args []string) {
	inFile, err := readTodoFile(conf.Todos)
	if err != nil {
		fmt.Println(err)
		return
	}
	var active todo.TaskList
	var archive todo.TaskList
	for _, task := range inFile {
		if task.Done {
			archive = append(archive, task)
		} else {
			active = append(active, task)
		}
	}

	if len(archive) == 0 {
		fmt.Println("No tasks archived.")
		return
	}

	arcFile, err := os.OpenFile(conf.Archive, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Couldn't archive:", err)
		return
	}
	defer arcFile.Close()
	for _, t := range archive {
		_, err := fmt.Fprintln(arcFile, t.Raw)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	file, err := os.Create(conf.Todos)
	if err != nil {
		fmt.Println("Couldn't modify todo file:", err)
		return
	}
	defer file.Close()
	for _, t := range active {
		_, err := fmt.Fprintln(file, t.Raw)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Printf("Archived %d task", len(archive))
	if len(archive) != 1 {
		fmt.Print("s")
	}
	fmt.Println(".")

}

func readTodoFile(f string) (todo.TaskList, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	todos, err := todo.FromReader(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return todos, nil
}
