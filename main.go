package main

import (
	"log"

	"gx/components" // Import the components package

	"github.com/gdamore/tcell/v2"
	"github.com/google/go-github/v65/github"
)

func main() {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defer s.Fini()

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)
	s.SetStyle(defStyle)

	searchTerm := ""
	cursorPosition := 0
	selectedButton := 0 // 0 for search button

	running := true
	showResults := false
	repos := []*github.Repository{}

	for running {
		s.Clear()

		// Draw search box and button
		components.DrawTextBox(s, 2, 2, 30, searchTerm, defStyle)
		components.DrawButton(s, 35, 2, 10, "Search", defStyle, selectedButton == 0)

		if showResults {
			components.DrawGitHubRepositories(s, 0, 5, 50, 10, repos, 0, 0, defStyle)
		}

		s.Show()

		// Poll for events
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				running = false
			case tcell.KeyEnter:
				// Submit search term when button is selected
				if selectedButton == 0 {
					client := getGitHubClient()
					var err error
					repos, err = searchRepositories(client, searchTerm)
					if err != nil {
						log.Fatalf("Error searching repositories: %v", err)
					}
					showResults = true
				}
			case tcell.KeyUp:
				// Move focus between buttons and text box
				if selectedButton > 0 {
					selectedButton--
				}
			case tcell.KeyDown:
				// Move focus between buttons and text box
				if selectedButton < 1 {
					selectedButton++
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				// Handle backspace in the search term
				if cursorPosition > 0 && len(searchTerm) > 0 {
					searchTerm = searchTerm[:len(searchTerm)-1]
					cursorPosition--
				}
			case tcell.KeyRune:
				// Handle text input into the search term box
				searchTerm += string(ev.Rune())
				cursorPosition++
			}
		case *tcell.EventResize:
			s.Sync()
		}
	}
}
