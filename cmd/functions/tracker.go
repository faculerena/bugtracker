package tracker

import (
	"encoding/json"
	"os"
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
