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
	"github.com/spf13/cobra"
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
		t.Print()

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

func (t *Tracker) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, bug := range *t {
		idx++
		task := bug.What
		done := "no"
		if bug.Solved {
			task = fmt.Sprintf("\u2705 %s", bug.What)
			done = "yes"
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: bug.Created.Format(time.RFC822)},
			{Text: bug.Completed.Format(time.RFC822)},
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
