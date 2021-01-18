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
	if t, ok = parse(args); !ok {
		return
	}
	var id int
	if id, ok = lists.Add(t); ok {
		fmt.Printf("Task %d created.\n", id)
	}
	return
}

// Claim moves a task from the backlog to the TODO list
func Claim(id int) (ok bool) {
	if ok = lists.Claim(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// Done marks a task as completed
func Done(id int) (ok bool) {
	if ok = lists.Done(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// Remove deletes a task entirely
func Remove(id int) (ok bool) {
	if ok = lists.Remove(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// Return moves a task from the TODO list back to the Backlog
func Return(id int) (ok bool) {
	if ok = lists.Return(id); !ok {
		fmt.Fprintf(os.Stderr, "Failed to find task '%d'\n", id)
	}
	return
}

// TODO prints the TODO list
func TODO() bool {
	return lists["todo"].Print()
}

// Backlog prints the Backlog
func Backlog() bool {
	return lists["backlog"].Print()
}

// Completed prints the Completed tasks
func Completed() bool {
	return lists["completed"].Print()
}

// All prints every tracked task
func All() (ok bool) {
	if ok = TODO(); !ok {
		return
	}
	fmt.Printf("\n\033[100m BACKLOG \033[49m\n\n")
	if ok = Backlog(); !ok {
		return
	}
	fmt.Printf("\n\033[100m COMPLETED \033[49m\n\n")
	ok = Completed()
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
