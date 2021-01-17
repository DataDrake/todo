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

package cli

import (
	"fmt"
	"github.com/DataDrake/cli-ng/cmd"
	"github.com/DataDrake/todo/colors"
	"github.com/DataDrake/todo/tasks"
	"os"
)

func init() {
	cmd.Register(&Color)
}

// Color a new task
var Color = cmd.Sub{
	Name:  "color",
	Short: "Set color for project or label",
	Args:  &ColorArgs{},
	Run:   ColorRun,
}

// ColorArgs accepts a color specification
type ColorArgs struct {
	Spec string `desc:"specification for a project or label (e.g. @project or :label "`
	Name string `desc:"name of color"`
}

// ColorRun carries out the "add" sub-command
func ColorRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*ColorArgs)
	if name, ok := tasks.ParseProject(args.Spec); ok {
		if ok := colors.Set("projects", name, args.Name); !ok {
			os.Exit(1)
		}
		return
	}
	if name, ok := tasks.ParseLabel(args.Spec); ok {
		if ok := colors.Set("labels", name, args.Name); !ok {
			os.Exit(1)
		}
		return
	}
	fmt.Fprintln(os.Stderr, "Spec must be for a project or label")
	os.Exit(1)
}
