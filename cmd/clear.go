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
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all bugs saved",
	Long: `You can use 'tracker clear' and a check message will appear, if you confirm it, the tracker 
will become empty as new.`,
	Run: func(cmd *cobra.Command, args []string) {
		confirmText := "DELETE ALL, PLEASE"
		fmt.Printf("Type %s to continue\n", confirmText)
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if input != confirmText {
			os.Exit(1)
		}

		err = os.Remove(".tracker.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		_, err = os.Create(".tracker.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("Please, remember to type 'tracker init' before adding a your first bug")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
