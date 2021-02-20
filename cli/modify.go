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
	cmd.Register(&Modify)
}

// Modify deletes a task permanently
var Modify = cmd.Sub{
	Name:  "modify",
	Alias: "mod",
	Short: "Modify an existing task",
	Args:  &ModifyArgs{},
	Run:   ModifyRun,
}

// ModifyArgs specifies the ID of the task to modify
type ModifyArgs struct {
	ID   uint     `desc:"ID of Task to modify"`
	Spec []string `desc:"Specification for modifying the task (e.g. @new_project :new_label \"New Name\")"`
}

// ModifyRun carries out the "modify" sub-command
func ModifyRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*ModifyArgs)
	if ok := tasks.Modify(args.ID, args.Spec); !ok {
		os.Exit(1)
	}
}
