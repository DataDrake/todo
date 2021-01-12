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
	"path/filepath"
)

var (
	todo      List
	backlog   List
	completed List
)

var (
	todoPath      string
	backlogPath   string
	completedPath string
)

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get homedir, reason: %s\n", err)
		os.Exit(1)
	}
	data := filepath.Join(home, ".local", "share", "todo")
	if err = os.MkdirAll(data, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to make directory '%s', reason: %s\n", data, err)
		os.Exit(1)
	}
	todoPath = filepath.Join(data, "todo.lst")
	backlogPath = filepath.Join(data, "backlog.lst")
	completedPath = filepath.Join(data, "completed.lst")
	if ok := loadAll(); !ok {
		os.Exit(1)
	}
}

// loadAll reads in all of the task lists
func loadAll() (ok bool) {
	if todo, ok = Load(todoPath); !ok {
		return
	}
	if backlog, ok = Load(backlogPath); !ok {
		return
	}
	completed, ok = Load(completedPath)
	return
}

// saveAll writes out all of the task lists
func saveAll() bool {
	return todo.Save(todoPath) && backlog.Save(backlogPath) && completed.Save(completedPath)
}

// nextID retrieves the next available ID, starting from 1
func nextID() (id int) {
	id = 1
	for {
		if index := todo.position(id); index != -1 {
			id++
			continue
		}
		if index := backlog.position(id); index != -1 {
			id++
			continue
		}
		if index := completed.position(id); index != -1 {
			id++
			continue
		}
		return
	}
}
