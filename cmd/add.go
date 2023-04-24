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
	"github.com/faculerena/bugtracker/internal"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a bug to the tracker",
	Long:  `tracker add initializes the ask interface to track a new bug`,
	Run: func(cmd *cobra.Command, args []string) {

		id, what, how, priority, _ := getInputAdd()

		t := &tracker.Bugs{}
		if err := t.Load(tracker.File); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		t.Add(id, what, how, priority, false)

		err := t.Store(tracker.File)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}
		fmt.Printf("Bug with ID %v created\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func getInputAdd() (int, string, string, int, bool) {

	idReturn, err := getNewId()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("What bug did you encounter?: ")
	fmt.Print("> ")
	whatInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	what := strings.TrimSuffix(whatInput, "\n")

	fmt.Println("How to reproduce it: ")
	fmt.Print("> ")
	howInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	how := strings.TrimSuffix(howInput, "\n")

	fmt.Println("What priority does it have?: ")
	fmt.Print("> ")
	priorityInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	priority := strings.TrimSuffix(priorityInput, "\n")
	priorityReturn, _ := strconv.Atoi(priority)

	return idReturn, what, how, priorityReturn, false

}

func getNewId() (int, error) {
	file, err := os.ReadFile(".id")
	if err != nil {
		fmt.Println("Please type 'tracker init', the id file doesn't exist")
		os.Exit(1)
	}

	idx, _ := strconv.Atoi(string(file))
	idreturn := idx
	idx++
	err = os.WriteFile(".id", []byte(strconv.Itoa(idx)), 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return idreturn, err

}
