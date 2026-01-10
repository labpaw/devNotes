package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Note represents a single note
type Note struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// App struct
type App struct {
	ctx       context.Context
	notesFile string
}

// NewApp creates a new App application struct
func NewApp() *App {
	homeDir, _ := os.UserHomeDir()
	notesDir := filepath.Join(homeDir, ".devnotes")
	os.MkdirAll(notesDir, 0755)
	
	return &App{
		notesFile: filepath.Join(notesDir, "notes.json"),
	}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetNotes returns all notes
func (a *App) GetNotes() []Note {
	notes := []Note{}
	data, err := os.ReadFile(a.notesFile)
	if err != nil {
		return notes
	}
	json.Unmarshal(data, &notes)
	return notes
}

// SaveNote creates or updates a note
func (a *App) SaveNote(id, content string) Note {
	notes := a.GetNotes()
	now := time.Now().Format(time.RFC3339)
	
	// Check if note exists (update)
	for i, note := range notes {
		if note.ID == id {
			notes[i].Content = content
			notes[i].UpdatedAt = now
			a.saveNotes(notes)
			return notes[i]
		}
	}
	
	// Create new note
	newNote := Note{
		ID:        generateID(),
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}
	notes = append([]Note{newNote}, notes...)
	a.saveNotes(notes)
	return newNote
}

// DeleteNote removes a note by ID
func (a *App) DeleteNote(id string) bool {
	notes := a.GetNotes()
	for i, note := range notes {
		if note.ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			a.saveNotes(notes)
			return true
		}
	}
	return false
}

// saveNotes writes notes to file
func (a *App) saveNotes(notes []Note) error {
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(a.notesFile, data, 0644)
}

// generateID creates a simple unique ID
func generateID() string {
	return time.Now().Format("20060102150405.000000")
}
