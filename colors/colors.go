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
	"sort"
	"strings"
)

// Names is a list of the supported colors
var Names = []string{
	"None",
	"Black",
	"DarkGray",
	"LightGray",
	"White",
	"Blue",
	"LightBlue",
	"Cyan",
	"LightCyan",
	"Yellow",
	"LightYellow",
	"Green",
	"LightGreen",
	"Magenta",
	"LightMagenta",
	"Red",
	"LightRed",
}

// Codes maps the name of a color to its escape code equivalent
var Codes = map[string]int{
	"None":         49,
	"Black":        40,
	"Red":          41,
	"Green":        42,
	"Yellow":       43,
	"Blue":         44,
	"Magenta":      45,
	"Cyan":         46,
	"LightGray":    47,
	"DarkGray":     100,
	"LightRed":     101,
	"LightGreen":   102,
	"LightYellow":  103,
	"LightBlue":    104,
	"LightMagenta": 105,
	"LightCyan":    106,
	"White":        107,
}

// Print deomonstrates all of the available colors on stdout
func Print() {
	for _, name := range Names {
		fmt.Printf("\033[%dm  \033[49m %s\n", Codes[name], name)
	}
	var keys []string
	for key := range configs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("\n\033[100m %s \033[49m\n\n", strings.ToUpper(key))
		for name, color := range configs[key] {
			fmt.Printf("\033[%dm  \033[49m %s\n", Codes[color], name)
		}
	}
}
