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
	"os"

	"github.com/spf13/cobra"
)

// listallCmd represents the listall command
var listallCmd = &cobra.Command{
	Use:   "listall ",
	Short: "List all bugs",
	Long: `use 'list' to retrieve ALL bugs saved on the tracker, or use 
'tracker [number] to retrieve the last [number] bugs saved'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		t := &trackfx.Tracker{}
		err := t.Load(trackfx.TrackerFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		t.ListAll()

	},
}

func init() {
	rootCmd.AddCommand(listallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
