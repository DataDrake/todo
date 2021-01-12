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

package colors

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	labelColors   Map
	projectColors Map
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
	labels := filepath.Join(data, "labels.lst")
	labelColors, err = load(labels)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load color config '%s', reason: %s\n", labels, err)
		os.Exit(1)
	}
	projects := filepath.Join(data, "projects.lst")
	projectColors, err = load(projects)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load color config '%s', reason: %s\n", projects, err)
		os.Exit(1)
	}
}

func load(path string) (m Map, err error) {
	m = make(Map)
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
			return
		}
		return
	}
	defer f.Close()
	for {
		name, color, err := read(f)
		if err != nil {
			break
		}
		m[name] = color
	}
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		err = fmt.Errorf("failed to read color, reason: %s", err)
		return
	}
	return
}

func saveAll() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get homedir, reason: %s", err)
	}
	data := filepath.Join(home, ".local", "share", "todo")
	labels := filepath.Join(data, "labels.lst")
	if err := labelColors.save(labels); err != nil {
		return fmt.Errorf("failed to save color config '%s', reason: %s", labels, err)
	}
	projects := filepath.Join(data, "projects.lst")
	if err := projectColors.save(projects); err != nil {
		return fmt.Errorf("failed to save color config '%s', reason: %s", projects, err)
	}
	return nil
}

// Map created associations between strings and color names
type Map map[string]string

func (m Map) save(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for name, color := range m {
		if err := write(f, name, color); err != nil {
			return err
		}
	}
	return nil
}

// LabelColor gets the color code for a label
func LabelColor(name string) (code int) {
	color, ok := labelColors[name]
	if !ok {
		color = "None"
	}
	code = Codes[color]
	return
}

// ProjectColor gets the color code for a project
func ProjectColor(name string) (code int) {
	color, ok := projectColors[name]
	if !ok {
		color = "None"
	}
	code = Codes[color]
	return
}
