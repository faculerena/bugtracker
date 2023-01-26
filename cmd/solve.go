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
	trackfx "github.com/faculerena/bugtracker/cmd/functions"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// solveCmd represents the solved command
var solveCmd = &cobra.Command{
	Use:   "solve <ID>",
	Short: "WIP", //"Mark a bug as solved",
	Long:  `Use 'tracker solved <ID>' to mark a bug as solved.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solved called")
		if len(args) != 1 {
			os.Exit(1)
		}
		t := &trackfx.Tracker{}
		if err := t.Load(trackfx.TrackerFile); err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		solvedID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}
		fmt.Println(solvedID)
		err = t.Solve(solvedID)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}
		err = t.Store(trackfx.TrackerFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}

	},
}

func init() {
	rootCmd.AddCommand(solveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// solveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// solveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}