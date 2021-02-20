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
	cmd.Register(&Claim)
}

// Claim moves a task from the Backlog to the TODO list
var Claim = cmd.Sub{
	Name:  "claim",
	Alias: "^",
	Short: "Move a task from the Backlog to TODO",
	Args:  &ClaimArgs{},
	Run:   ClaimRun,
}

// ClaimArgs specifies the ID of the task to claim
type ClaimArgs struct {
	IDs []uint `desc:"ID(s) of Task(s) to claim as TODO"`
}

// ClaimRun carries out the "claim" sub-command
func ClaimRun(r *cmd.Root, s *cmd.Sub) {
	args := s.Args.(*ClaimArgs)
	for _, id := range args.IDs {
		if ok := tasks.Claim(id); !ok {
			os.Exit(1)
		}
	}
}
