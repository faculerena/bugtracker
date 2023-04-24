/*
Copyright © 2023 Facundo Lerena  <contacto@faculerena.com.ar>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"github.com/faculerena/bugtracker/internal"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// reopenCmd represents the reopen command
var reopenCmd = &cobra.Command{
	Use:   "reopen <ID>",
	Short: "WIP", //"Mark a solved bug as open again",
	Long: `Use 'reopen <ID>' to mark a solved bug as open again. If you put a still 
open bug, a message will appear and nothing will be modified.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			os.Exit(1)
		}
		t := &tracker.Bugs{}

		if err := t.Load(tracker.File); err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		solvedID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}
		fmt.Println(solvedID)
		err = t.Reopen(solvedID)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}

		err = t.Store(tracker.File)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}

		fmt.Printf("Reopened bug %v\n", solvedID)

	},
}

func init() {
	rootCmd.AddCommand(reopenCmd)
}
