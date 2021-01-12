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
	"os"
	"time"
)

// Add creates a new Task and places it in the backlog
func Add(args []string) (ok bool) {
	var t Task
	if t, ok = parse(args); !ok {
		return
	}
	t.ID = nextID()
	t.Created = time.Now()
	backlog = append(backlog, t)
	if ok = saveAll(); !ok {
		return
	}
	fmt.Printf("Task %d created.\n", t.ID)
	return
}

// Claim moves a task from the backlog to the TODO list
func Claim(id int) bool {
	if l, t, ok := backlog.remove(id); ok {
		backlog = l
		todo = append(todo, t)
		return saveAll()
	}
	fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	return false
}

// Done marks a task as completed
func Done(id int) bool {
	if l, t, ok := todo.remove(id); ok {
		todo = l
		t.Finished = time.Now()
		completed = append(completed, t)
		return saveAll()
	}
	if l, t, ok := backlog.remove(id); ok {
		backlog = l
		t.Finished = time.Now()
		completed = append(completed, t)
		return saveAll()
	}
	fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	return false
}

// Remove deletes a task entirely
func Remove(id int) bool {
	if l, _, ok := todo.remove(id); ok {
		todo = l
		return saveAll()
	}
	if l, _, ok := backlog.remove(id); ok {
		backlog = l
		return saveAll()
	}
	if l, _, ok := completed.remove(id); ok {
		completed = l
		return saveAll()
	}
	fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	return false
}

// Return moves a task from the TODO list back to the Backlog
func Return(id int) bool {
	if l, t, ok := todo.remove(id); ok {
		todo = l
		backlog = append(backlog, t)
		return saveAll()
	}
	fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	return false
}

// TODO prints the TODO list
func TODO() bool {
	return todo.Print()
}

// Backlog prints the Backlog
func Backlog() bool {
	return backlog.Print()
}

// Completed prints the Completed tasks
func Completed() bool {
	return completed.Print()
}

// All prints every tracked task
func All() (ok bool) {
	if ok = todo.Print(); !ok {
		return
	}
	fmt.Printf("\nBacklog:\n\n")
	if ok = backlog.Print(); !ok {
		return
	}
	fmt.Printf("\nCompleted:\n\n")
	ok = completed.Print()
	return
}

// ResetCompleted permanently deletes all completed tasks
func ResetCompleted() bool {
	completed = make(List, 0)
	return saveAll()
}
