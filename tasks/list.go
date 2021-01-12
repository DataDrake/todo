//
// Copyright 2021 Bryan T. Meyers <root@datadrake.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tasks

import (
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
)

// List is a list of one or more tasks
type List []Task

// Len gets the length of a List for sorting
func (l List) Len() int {
	return len(l)
}

// Less compares Tasks for sorting: Project -> Label -> ID
func (l List) Less(i, j int) bool {
	if l[i].Project < l[j].Project {
		return true
	}
	if l[i].Project > l[j].Project {
		return false
	}
	if l[i].Label < l[j].Label {
		return true
	}
	if l[i].Label > l[j].Label {
		return false
	}
	return l[i].ID < l[j].ID
}

// Swap exchanges elements for sorting
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// Load reads in all of the Tasks in a file, failing silently if the file does not exist
func Load(path string) (l List, ok bool) {
	r, err := os.Open(path)
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stdout, "Failed to load tasks from '%s', reason: %s\n", path, err)
			return
		}
		ok = true
		return
	}
	defer r.Close()
	var t Task
	for {
		t, err = Read(r)
		if err != nil {
			break
		}
		l = append(l, t)
	}
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		fmt.Fprintf(os.Stderr, "Failed to read task, reason: %s\n", err)
		return
	}
	sort.Sort(l)
	ok = true
	return
}

// Save writes a List out to file
func (l List) Save(path string) (ok bool) {
	w, err := os.Create(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open '%s', reason: %s\n", path, err)
		return
	}
	defer w.Close()
	sort.Sort(l)
	for _, t := range l {
		if err = t.Write(w); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write task, reason: %s\n", err)
			return
		}
	}
	ok = true
	return
}

// Print writes a List out to console
func (l List) Print() bool {
	if len(l) == 0 {
		fmt.Println("No tasks yet.")
		return true
	}
	sort.Sort(l)
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()
	if _, err := fmt.Fprintln(tw, "\033[1mID\tCreated\tFinished\t\033[49m \033[49m\tProject\t\033[49m \033[49m\tLabel\tName\033[0m"); err != nil {
		fmt.Fprintf(os.Stderr, "Error while printing: %s\n", err)
		return false
	}
	for _, t := range l {
		if err := t.Print(tw); err != nil {
			fmt.Fprintf(os.Stderr, "Error while printing: %s\n", err)
			return false
		}
	}
	return true
}

// position finds the index of the Task with a matching ID or -1 if not found
func (l List) position(id int) int {
	for i, t := range l {
		if t.ID == id {
			return i
		}
	}
	return -1
}

// Find searches for a task by ID, returning it and its index if found
func (l List) find(id int) (t Task, index int, ok bool) {
	if index = l.position(id); index == -1 {
		return
	}
	t = l[index]
	ok = true
	return
}

// Remove deletes a task from the list by ID
func (l List) remove(id int) (rest List, t Task, ok bool) {
	t, i, ok := l.find(id)
	if !ok {
		return
	}
	if i == len(l)-1 {
		rest = l[:i]
		return
	}
	rest = append(l[:i], l[i+1:]...)
	return
}
