package tracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"os"
	"strconv"
	"time"
)

/* A bug consists of:
ID
what happened string
how it happened string
priority 1 (top) - 5 (low) int
solved bool
time created
time solved
*/

const (
	TrackerFile = ".tracker.json"
	NotesFile   = ".notes.json"
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

func (t *Tracker) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0664)

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

func (t *Tracker) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
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
			{Align: simpletable.AlignCenter, Text: "Solved"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, bug := range *t {
		id := bug.ID
		what := bug.What
		how := bug.Steps
		priority := bug.Priority
		created := bug.Created

		solved := bug.Solved
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", id)},
			{Text: what},
			{Text: how},
			{Text: strconv.Itoa(priority)},
			{Text: created.Format(time.RFC822)},
			{Text: strconv.FormatBool(solved)},
		})

	}

	table.Body = &simpletable.Body{Cells: cells}
	/*
		table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("You have %d pending todos", t.CountPending()))},
		}}
	*/
	table.SetStyle(simpletable.StyleRounded)

	table.Println()
}

func (t *Tracker) Complete(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].Completed = time.Now()
	ls[index-1].Solved = true

	return nil
}
