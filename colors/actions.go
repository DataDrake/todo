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
	"os"
)

// SetLabel sets the color for a Label
func SetLabel(name, color string) bool {
	if _, ok := Codes[color]; !ok {
		fmt.Fprintf(os.Stderr, "Color '%s' does not exist.\n", color)
		return false
	}
	if color == "None" {
		delete(labelColors, name)
	} else {
		labelColors[name] = color
	}
	return Save()
}

// SetProject sets the color for a project
func SetProject(name, color string) bool {
	if _, ok := Codes[color]; !ok {
		fmt.Fprintf(os.Stderr, "Color '%s' does not exist.\n", color)
		return false
	}
	if color == "None" {
		delete(projectColors, name)
	} else {
		projectColors[name] = color
	}
	return Save()
}

// Save writes out color configs to disk
func Save() bool {
	if err := saveAll(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save colors, reason: %s\n", err)
		return false
	}
	return true
}
