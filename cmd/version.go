// Copyright © 2017 Ahmed T. Ali <ah.tajelsir@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	// Native
	"fmt"

	// Packages
	"github.com/spf13/cobra"
)

// VERSION will be automatically set by GoReleaser to the current Git tag
// (the 'v' prefix is stripped)
var VERSION = "X.Y.Z"

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"about"},
	Args:    cobra.MaximumNArgs(0),
	Short:   "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
  ╭────────────────────────────────────────────╮
  │   /    Git todos ( v%s )                │
  │  /|_                                       │
  │ /_ /   By Ahmed T. Ali (https://ahmed.sd)  │
  │   /                                        │
  │  /     Happy Coding!                       │
  ╰────────────────────────────────────────────╯

`, VERSION)
	},
}

func init() {
	appCmd.AddCommand(versionCmd)
}
