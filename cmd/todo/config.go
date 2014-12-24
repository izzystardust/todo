// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	Todos   string
	Archive string
}

func GetConfiguration() config {
	home := os.Getenv("HOME")
	file, err := os.Open(home + "/.config/todo/todo.toml")
	defaultConf := config{
		Todos:   home + "/todo.txt",
		Archive: home + "/.config/todo/archive.txt",
	}

	if err != nil {
		// no config, use defaults
		return defaultConf
	}

	var a config
	_, err = toml.DecodeReader(file, &a)
	if err != nil {
		fmt.Println(err)
		return defaultConf
	}
	a.Todos = os.ExpandEnv(a.Todos)
	a.Archive = os.ExpandEnv(a.Archive)
	return a
}
