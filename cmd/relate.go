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

	"github.com/spf13/cobra"
)

// relateCmd represents the relate command
var relateCmd = &cobra.Command{
	Use:   "relate <ID> <ID_target>",
	Short: "WIP", //Relate a bug with another
	Long: `Use relate <ID_to_relate> <ID_target_bug> to add the first bug to the "related bugs" 
in the target bug.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("relate called")
	},
}

func init() {
	rootCmd.AddCommand(relateCmd)
}
