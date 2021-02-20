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
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/DataDrake/todo/tasks"
	"os"
)

func init() {
	cmd.Register(&List)
}

// List prints the TODO list
var List = cmd.Sub{
	Name:  "list",
	Alias: "ls",
	Short: "Print current TODO list",
	Flags: &ListFlags{},
	Run:   ListRun,
}

// ListFlags contains the flags for the "list" sub-command
type ListFlags struct {
	Project  string `short:"p" long:"project" desc:"Name of project to filter on"`
	Label    string `short:"l" long:"label" desc:"Name of label to fiter on"`
	FullTime bool   `short:"T" long:"full-time" desc:"Print both date and time for all tasks"`
}

// ListRun carries out the "list" sub-command
func ListRun(r *cmd.Root, s *cmd.Sub) {
	flags := s.Flags.(*ListFlags)
	if ok := tasks.TODO(flags.Project, flags.Label, flags.FullTime); !ok {
		os.Exit(1)
	}
}
