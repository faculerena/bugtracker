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
	tracker "github.com/faculerena/bugtracker/cmd/tracker"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <ID>",
	Short: "WIP", //"Remove a bug from the tracker",
	Long:  `If you want to remove a bug, you can use 'remove <ID>' to delete it from the tracker.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
		if len(args) != 1 {
			os.Exit(1)
		}
		t := &tracker.Bugs{}
		if err := t.Load(tracker.File); err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		deletingID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}
		fmt.Println(deletingID)
		err = t.Remove(deletingID)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}
		err = t.Store(tracker.File)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}

	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
