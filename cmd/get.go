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
	tracker "github.com/faculerena/bugtracker/internal"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")

		if len(args) != 1 {
			os.Exit(1)
		}
		t := &tracker.Bugs{}

		if err := t.Load(tracker.File); err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}

		editID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}

		err = t.Get(editID)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(4)
		}

		err = t.Store(tracker.File)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(5)
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
