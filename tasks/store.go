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
	"github.com/DataDrake/todo/store"
	"os"
	"strings"
	"time"
)

var lists Store

func init() {
	lists = make(Store)
	if ok := lists.load(store.Path()); !ok {
		os.Exit(1)
	}
}

// Store contains multiple lists of tasks
type Store map[string]List

var storeLists = []string{"backlog", "completed", "todo"}

func (s Store) load(data string) (ok bool) {
	for _, name := range storeLists {
		if s[name], ok = load(data, name); !ok {
			return
		}
	}
	return
}

// saveAll writes out all of the task lists
func (s Store) save(data string) (ok bool) {
	for name, list := range s {
		if ok = list.Save(data, name); !ok {
			return
		}
	}
	return
}

// nextID retrieves the next available ID, starting from 1
func (s Store) nextID() (id uint) {
	id = 1
	for {
		found := false
		for _, list := range s {
			if index := list.position(id); index != -1 {
				found = true
				break
			}
		}
		if !found {
			return
		}
		id++
	}
}

// Add creates a new Task and places it in the backlog
func (s Store) Add(t Task) (id uint, ok bool) {
	id = s.nextID()
	t.ID = id
	t.Created = time.Now()
	s["backlog"] = append(s["backlog"], t)
	ok = s.save(store.Path())
	return
}

// Claim moves a task from the backlog to the TODO list
func (s Store) Claim(id uint) (ok bool) {
	if l, t, ok := s["backlog"].remove(id); ok {
		s["backlog"] = l
		s["todo"] = append(s["todo"], t)
		return s.save(store.Path())
	}
	return
}

// Done marks a task as completed
func (s Store) Done(id uint) bool {
	for _, name := range []string{"todo", "backlog"} {
		if l, t, ok := s[name].remove(id); ok {
			s[name] = l
			t.Finished = time.Now()
			s["completed"] = append(s["completed"], t)
			return s.save(store.Path())
		}
	}
	return false
}

// Modify updates an existing Task
func (s Store) Modify(t Task) (ok bool) {
	for name, list := range s {
		if l, ok := list.modify(t); ok {
			s[name] = l
			return s.save(store.Path())
		}
	}
	return false
}

// Remove deletes a task entirely
func (s Store) Remove(id uint) bool {
	for name, list := range s {
		if l, _, ok := list.remove(id); ok {
			s[name] = l
			return s.save(store.Path())
		}
	}
	return false
}

// Return moves a task from the TODO list back to the Backlog
func (s Store) Return(id uint) bool {
	if l, t, ok := s["todo"].remove(id); ok {
		s["todo"] = l
		s["backlog"] = append(s["backlog"], t)
		return s.save(store.Path())
	}
	return false
}

// Reset permanently deletes all tasks in a list
func (s Store) Reset(name string) bool {
	s[name] = make(List, 0)
	return s.save(store.Path())
}

// Report generates a TODO.md for the current set of tasks
func (s Store) Report() (ok bool) {
	r, err := os.Create("TODO.md")
	if err != nil {
		fmt.Printf("Failed to open 'TODO.md', reason: %s\n", err)
		return
	}
	defer r.Close()
	for _, name := range []string{"todo", "backlog", "completed"} {
		fmt.Fprintf(r, "# %s\n\n", strings.ToUpper(name))
		if err = s[name].Report(r); err != nil {
			fmt.Printf("Problem writing report for '%s': %s\n", name, err)
			return
		}
		fmt.Fprintln(r)
	}
	ok = true
	return
}

// Undo moves a Task from Completed to the Backlog and resets its completion time
func (s Store) Undo(id uint) (ok bool) {
	if l, t, ok := s["completed"].remove(id); ok {
		s["completed"] = l
		t.Finished = time.Time{}
		s["backlog"] = append(s["backlog"], t)
		return s.save(store.Path())
	}
	return
}
