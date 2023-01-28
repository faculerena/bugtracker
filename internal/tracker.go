package tracker

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"os"
	"strconv"
	"strings"
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
	File = ".tracker.json"
	//NotesFile = ".notes.json"
)

type Bug struct {
	ID        int
	What      string
	Steps     string
	Priority  int
	Created   time.Time
	Completed time.Time
	Solved    bool
	Related   []int
}

type Bugs []Bug

func (t *Bugs) Load(filename string) error {
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

func (t *Bugs) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0664)

}

func (t *Bugs) Add(ID int, what string, steps string, priority int, solved bool) {
	var related []int
	tracker := Bug{
		ID,
		what,
		steps,
		priority,
		time.Now(),
		time.Time{},
		solved,
		related,
	}
	*t = append(*t, tracker)
}

func (t *Bugs) List() {

	_, err := os.Open(File)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Tracker not found. Please, type 'tracker init'.")
		os.Exit(1)
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "What?"},
			{Align: simpletable.AlignCenter, Text: "How?"},
			{Align: simpletable.AlignCenter, Text: "Priority"},
			{Align: simpletable.AlignCenter, Text: "Created"},
			{Align: simpletable.AlignCenter, Text: "Solved"},
			{Align: simpletable.AlignCenter, Text: "Related"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, bug := range *t {
		id := bug.ID
		what := bug.What
		how := bug.Steps
		priority := bug.Priority
		created := bug.Created
		solved := Red("No")
		related := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(bug.Related)), ","), "[]")

		if bug.Related == nil {
			related = ("none")
		}

		if bug.Solved == true {
			continue
		}
		if id == -1 {
			continue
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", id)},
			{Text: what},
			{Text: how},
			{Text: strconv.Itoa(priority)},
			{Text: created.Format(time.RFC822)},
			{Text: solved},
			{Text: related},
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

func (t *Bugs) ListAll() {

	_, err := os.Open(File)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Tracker not found. Please, type 'tracker init'.")
		os.Exit(1)
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "What?"},
			{Align: simpletable.AlignCenter, Text: "How?"},
			{Align: simpletable.AlignCenter, Text: "Priority"},
			{Align: simpletable.AlignCenter, Text: "Created"},
			{Align: simpletable.AlignCenter, Text: "Solved"},
			{Align: simpletable.AlignCenter, Text: "Related"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, bug := range *t {
		id := bug.ID
		what := bug.What
		how := bug.Steps
		priority := bug.Priority
		created := bug.Created
		solved := Red("No")
		related := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(bug.Related)), ","), "[]")

		if bug.Related == nil {
			related = ("none")
		}

		if bug.Solved == true {
			solved = Green("Yes")

		}

		if id == -1 {
			continue
		}

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", id)},
			{Text: what},
			{Text: how},
			{Text: strconv.Itoa(priority)},
			{Text: created.Format(time.RFC822)},
			{Text: solved},
			{Text: related},
		})

	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleRounded)

	table.Println()
}

func (t *Bugs) Solve(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].Completed = time.Now()
	ls[index-1].Solved = true

	return nil
}

func (t *Bugs) Remove(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	ls[index-1].ID = -1
	ls[index-1].Solved = false
	ls[index-1].Completed = time.Now()
	ls[index-1].Created = time.Now()
	ls[index-1].What = ""
	ls[index-1].Steps = ""
	ls[index-1].Priority = 0

	return nil
}

func (t *Bugs) Reopen(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].Completed = time.Time{}
	ls[index-1].Solved = false

	return nil
}

func (t *Bugs) Relate(id int) error {
	ls := *t
	file, err := os.ReadFile(".id")
	if err != nil {
		os.Exit(1)
		return err
	}
	nextID, _ := strconv.Atoi(string(file))
	if id >= nextID {
		os.Exit(1)
		return err
	}

	ls[id-1].Related = append(ls[id-1].Related, id)

	return nil
}
