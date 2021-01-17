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
	"github.com/DataDrake/todo/store"
	"os"
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
func (s Store) nextID() (id int) {
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
func (s Store) Add(t Task) (id int, ok bool) {
	id = s.nextID()
	t.ID = id
	t.Created = time.Now()
	s["backlog"] = append(s["backlog"], t)
	ok = s.save(store.Path())
	return
}

// Claim moves a task from the backlog to the TODO list
func (s Store) Claim(id int) (ok bool) {
	var t Task
	if s["backlog"], t, ok = s["backlog"].remove(id); ok {
		s["todo"] = append(s["todo"], t)
		return s.save(store.Path())
	}
	return
}

// Done marks a task as completed
func (s Store) Done(id int) (ok bool) {
	var t Task
	for _, name := range []string{"todo", "backlog"} {
		if s[name], t, ok = s[name].remove(id); ok {
			t.Finished = time.Now()
			s["completed"] = append(s["completed"], t)
			return s.save(store.Path())
		}
	}
	return
}

// Remove deletes a task entirely
func (s Store) Remove(id int) (ok bool) {
	for name, list := range s {
		if s[name], _, ok = list.remove(id); ok {
			return s.save(store.Path())
		}
	}
	return
}

// Return moves a task from the TODO list back to the Backlog
func (s Store) Return(id int) (ok bool) {
	var t Task
	if s["todo"], t, ok = s["todo"].remove(id); ok {
		s["backlog"] = append(s["backlog"], t)
		return s.save(store.Path())
	}
	return
}

// Reset permanently deletes all tasks in a list
func (s Store) Reset(name string) bool {
	s[name] = make(List, 0)
	return s.save(store.Path())
}
