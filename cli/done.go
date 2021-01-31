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
	cmd.Register(&Done)
}

// Done marks a task as completed
var Done = cmd.Sub{
	Name:  "done",
	Alias: "!",
	Short: "Mark task done",
	Args:  &DoneArgs{},
	Run:   DoneRun,
}

// DoneArgs specifies the ID of the task to mark done
type DoneArgs struct {
	ID uint64 `desc:"ID of Task to mark as done"`
}

// DoneRun carries out the "done" sub-command
func DoneRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*DoneArgs)
	if ok := tasks.Done(int(args.ID)); !ok {
		os.Exit(1)
	}
}
