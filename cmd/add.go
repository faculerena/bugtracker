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
	"encoding/json"
	"fmt"
	trackfx "github.com/faculerena/bugtracker/cmd/functions"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"time"
)

type Bug struct {
	ID        int
	What      string
	Steps     string
	Priority  int
	Created   time.Time
	Completed time.Time
	Solved    bool
}

type Tracker []Bug

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a bug to the tracker",
	Long:  `tracker add initializes the ask interface to track a new bug`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		id, what, how, priority, _ := getInputAdd()
		fmt.Println(id, what, how, priority)
		t := &Tracker{}
		t.Add(id, what, how, priority, false)
		err := t.Store(trackfx.TrackerFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getInputAdd() (int, string, string, int, bool) {

	fmt.Println("ID?: ")
	idInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	id := strings.TrimSuffix(idInput, "\n")
	idReturn, _ := strconv.Atoi(id)

	fmt.Println("What bug did you encounter?: ")
	whatInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	what := strings.TrimSuffix(whatInput, "\n")

	fmt.Println("How to reproduce it: ")
	howInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	how := strings.TrimSuffix(howInput, "\n")

	fmt.Println("What priority does it have?: ")
	priorityInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	priority := strings.TrimSuffix(priorityInput, "\n")
	priorityReturn, _ := strconv.Atoi(priority)

	return idReturn, what, how, priorityReturn, false

}

func (t *Tracker) Add(ID int, what string, steps string, priority int, solved bool) {
	tracker := Bug{
		ID,
		what,
		steps,
		priority,
		time.Now(),
		time.Time{},
		solved,
	}
	*t = append(*t, tracker)
}

func (t *Tracker) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0664)

}
