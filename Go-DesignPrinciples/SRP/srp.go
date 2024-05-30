package SRP

import (
	"log"
	"os"
	"strings"
)

var entriesCount int

type Journal struct {
	Entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.Entries, "\n")
}

func (j *Journal) AddEntries(text string) int {
	entriesCount++
	j.Entries = append(j.Entries, text)
	return entriesCount
}

func (j *Journal) RemoveEntries(text string) int {
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

//Now code which breaks SPR

func (j *Journal) Save(filename string) {
	if err := os.WriteFile(filename, []byte(j.String()), 0644); err != nil {
		log.Panic("Error Occured", err)
	}
}

//Simlar like Load() , LoadFromWeb()

// Handling SPR
type Persistance struct {
	LineSeperator string
}

func (p *Persistance) SaveToFile(j *Journal, filename string) {
	_ = os.WriteFile(filename, []byte(strings.Join(j.Entries, p.LineSeperator)), 0644)
}
