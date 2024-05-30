package SRP

import (
	"log"
	"os"
	"strings"
)

// Journal struct responsible for managing journal entries
type Journal2 struct {
	Entries []string
}

func (j *Journal) String2() string {
	return strings.Join(j.Entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	j.Entries = append(j.Entries, text)
	return len(j.Entries)
}

func (j *Journal) RemoveEntry(text string) int {
	for i := 0; i < len(j.Entries); {
		if text == j.Entries[i] {
			if i == len(j.Entries)-1 {
				j.Entries = j.Entries[:i]
			} else {
				j.Entries = append(j.Entries[:i], j.Entries[i+1:]...)
			}
		} else {
			i++
		}
	}
	return len(j.Entries)
}

// Persistence struct responsible for saving to file
type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	if err := os.WriteFile(filename, []byte(strings.Join(j.Entries, p.LineSeparator)), 0644); err != nil {
		log.Panic("Error Occurred", err)
	}
}

// Application struct that composes Journal and Persistence
type Application struct {
	Journal     *Journal
	Persistence *Persistence
}

func main() {
	journal := &Journal{}
	persistence := &Persistence{LineSeparator: "\n"}

	app := &Application{
		Journal:     journal,
		Persistence: persistence,
	}

	app.Journal.AddEntry("First entry")
	app.Journal.AddEntry("Second entry")

	app.Persistence.SaveToFile(app.Journal, "journal.txt")
}
