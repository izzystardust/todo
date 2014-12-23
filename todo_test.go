// Copyright 2014 Ethan Miller. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package todo

import "testing"

func TestString(t *testing.T) {
	todos := []struct {
		todo   Task
		expect string
	}{
		{
			Task{Title: "Hello"}, "Hello",
		},
	}

	for _, i := range todos {
		if i.todo.String() != i.expect {
			t.Errorf("Got %v, expected %v", i.todo.String(), i.expect)
		}
	}
}

func TestParse(t *testing.T) {
	errors := []struct {
		in     string
		expect string
	}{
		{in: "x", expect: "todo: contains only done marker"},
		{in: "x 2014-11-12", expect: "todo: contains only done marker and completion time"},
	}

	good := []struct {
		in     string
		expect string
	}{
		{"Hello", "Hello"},
		{"x Hello", "Hello"},
	}

	for _, cas := range errors {
		_, e := Parse(cas.in)
		if e == nil {
			t.Errorf("On case %v, got no error (expected %v)", cas.in, cas.expect)
		} else if e.Error() != cas.expect {
			t.Errorf("On case %v, got %v (expected %v)", cas.in, e.Error(), cas.expect)
		}
	}

	for _, cas := range good {
		todo, e := Parse(cas.in)
		if e != nil {
			t.Errorf("On case %v, unexpected parse error %v", cas.in, e.Error())
		}
		if todo.String() != cas.expect {
			t.Errorf("On case %v, got %v (expected %v)", cas.in, todo.String(), cas.expect)
		}
	}
}
