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
	"github.com/DataDrake/todo/store"
	"io"
	"os"
	"path/filepath"
)

var configs map[string]Config

func init() {
	data := store.Path()
	configs = make(map[string]Config)
	for _, name := range []string{"labels", "projects"} {
		if err := load(data, name); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to load color config '%s', reason: %s\n", name, err)
			os.Exit(1)
		}
	}
}

func load(data, name string) (err error) {
	path := filepath.Join(data, name+".lst")
	c := make(Config)
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
			configs[name] = c
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
		c[name] = color
	}
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		err = fmt.Errorf("failed to read color, reason: %s", err)
		return
	}
	configs[name] = c
	return
}

func saveAll() error {
	data := store.Path()
	for name, config := range configs {
		if err := config.save(data, name); err != nil {
			return fmt.Errorf("failed to save color config '%s', reason: %s", name, err)
		}
	}
	return nil
}

// Config created associations between strings and color names
type Config map[string]string

func (c Config) save(data, name string) error {
	path := filepath.Join(data, name+".lst")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	for name, color := range c {
		if err := write(f, name, color); err != nil {
			return err
		}
	}
	return nil
}

// Color gets the color code from a color config
func Color(config, name string) (code int) {
	m, ok := configs[config]
	var color string
	if ok {
		color, ok = m[name]
	}
	if !ok {
		color = "None"
	}
	code = Codes[color]
	return
}
