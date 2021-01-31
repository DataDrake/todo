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
	cmd.Register(&Remove)
}

// Remove deletes a task permanently
var Remove = cmd.Sub{
	Name:  "remove",
	Short: "Remove a task entirely",
	Args:  &RemoveArgs{},
	Run:   RemoveRun,
}

// RemoveArgs specifies the ID of the task to remove
type RemoveArgs struct {
	ID uint64 `desc:"ID of Task to remove"`
}

// RemoveRun carries out the "remove" sub-command
func RemoveRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*RemoveArgs)
	if ok := tasks.Remove(int(args.ID)); !ok {
		os.Exit(1)
	}
}
