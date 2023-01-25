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
	"github.com/alexeyco/simpletable"
	trackfx "github.com/faculerena/bugtracker/cmd/functions"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [number]",
	Short: "List all bugs, or the last n bugs",
	Long: `use 'list' to retrieve ALL bugs saved on the tracker, or use 
'tracker [number] to retrieve the last [number] bugs saved'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		t := &Tracker{}
		err := t.Load(trackfx.TrackerFile)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		t.List()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (t *Tracker) List() {

	table := simpletable.New()

	/*
		trackfx.Bug{
			ID:        0,
			What:      "",
			Steps:     "",
			Priority:  0,
			Created:   time.Time{},
			Completed: time.Time{},
			Solved:    false,
		}

	*/

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "What?"},
			{Align: simpletable.AlignCenter, Text: "How?"},
			{Align: simpletable.AlignCenter, Text: "Priority"},
			{Align: simpletable.AlignCenter, Text: "Created"},
			{Align: simpletable.AlignCenter, Text: "Completed"},
			{Align: simpletable.AlignCenter, Text: "Solved"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, bug := range *t {
		id := bug.ID
		what := bug.What
		how := bug.Steps
		priority := bug.Priority
		created, completed := bug.Created, bug.Completed
		solved := bug.Solved

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", id)},
			{Text: what},
			{Text: how},
			{Text: strconv.Itoa(priority)},
			{Text: created.Format(time.RFC822)},
			{Text: completed.Format(time.RFC822)},
			{Text: strconv.FormatBool(solved)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	/*
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("You have %d pending todos", t.CountPending()))},
		}}
	*/
	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
