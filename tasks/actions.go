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
)

// Add creates a new Task and places it in the backlog
func Add(args []string) (ok bool) {
	var t Task
	if t, ok = parse(args, false); !ok {
		return
	}
	var id uint
	if id, ok = lists.Add(t); ok {
		fmt.Printf("Task %d created.\n", id)
	}
	return
}

// Claim moves a task from the backlog to the TODO list
func Claim(id uint) (ok bool) {
	if ok = lists.Claim(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// Done marks a task as completed
func Done(id uint) (ok bool) {
	if ok = lists.Done(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// Modify alters an existing task
func Modify(id uint, args []string) (ok bool) {
	var t Task
	if t, ok = parse(args, true); !ok {
		return
	}
	t.ID = id
	if ok = lists.Modify(t); ok {
		fmt.Printf("Task %d modified.\n", id)
	}
	return
}

// Remove deletes a task entirely
func Remove(id uint) (ok bool) {
	if ok = lists.Remove(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// Return moves a task from the TODO list back to the Backlog
func Return(id uint) (ok bool) {
	if ok = lists.Return(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// Undo moves a task from Completed to the Backlog
func Undo(id uint) (ok bool) {
	if ok = lists.Undo(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// TODO prints the TODO list
func TODO(project, label string, fullTime bool) bool {
	return lists["todo"].Print(project, label, fullTime)
}

// Backlog prints the Backlog
func Backlog(project, label string, fullTime bool) bool {
	return lists["backlog"].Print(project, label, fullTime)
}

// Completed prints the Completed tasks
func Completed(project, label string, fullTime bool) bool {
	return lists["completed"].Print(project, label, fullTime)
}

// All prints every tracked task
func All(project, label string, fullTime bool) (ok bool) {
	if ok = TODO(project, label, fullTime); !ok {
		return
	}
	fmt.Printf("\n\033[100m BACKLOG \033[49m\n\n")
	if ok = Backlog(project, label, fullTime); !ok {
		return
	}
	fmt.Printf("\n\033[100m COMPLETED \033[49m\n\n")
	ok = Completed(project, label, fullTime)
	return
}

// Report generates a TODO.md for the current set of tasks
func Report() bool {
	return lists.Report()
}

// ResetCompleted permanently deletes all completed tasks
func ResetCompleted() bool {
	return lists.Reset("completed")
}
