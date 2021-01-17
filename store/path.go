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

package store

import (
	"fmt"
	"os"
	"path/filepath"
)

var path string

// Path returns the directory path for all of the todo data
func Path() string {
	if len(path) > 0 {
		return path
	}
	if _, err := os.Stat(".todo"); err == nil {
		path = ".todo"
		return path
	}
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get homedir, reason: %s\n", err)
		os.Exit(1)
	}
	path = filepath.Join(home, ".local", "share", "todo")
	if err = os.MkdirAll(path, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to make directory '%s', reason: %s\n", path, err)
		os.Exit(1)
	}
	return path
}
