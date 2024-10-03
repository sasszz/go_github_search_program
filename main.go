package main

import (
	"log"
	"time"

	"gx/components"

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
	startIndex := 0
	focusedElement := 0 // 0 for search box, 1 for search button, 2 for results list
	displayHeight := 10
	loading := false
	showResults := false
	repos := []*github.Repository{}

	for running := true; running; {
		s.Clear()

		components.DrawTextBox(s, 2, 2, 50, searchTerm, defStyle)
		components.DrawButton(s, 55, 2, 10, "Search", defStyle, focusedElement == 1)

		if loading {
			components.DrawText(s, 2, 5, "Data loading... This can take up to 30 seconds.", defStyle)
		} else if showResults {
			components.DrawGitHubRepositories(s, 0, 5, 70, displayHeight, repos, startIndex, cursorPosition, defStyle)
		}

		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyCtrlC:
				running = false
			case tcell.KeyTab:
				if showResults {
					focusedElement = (focusedElement + 1) % 3
				} else {
					focusedElement = (focusedElement + 1) % 2
				}
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				if focusedElement == 0 && len(searchTerm) > 0 {
					searchTerm = searchTerm[:len(searchTerm)-1]
				}
			case tcell.KeyRune:
				if focusedElement == 0 {
					searchTerm += string(ev.Rune())
				}
			case tcell.KeyEnter:
				if focusedElement == 1 {
					loading = true
					s.Sync() // Force redraw immediately after entering "Search"
					go func() {
						time.Sleep(2 * time.Second) // Simulate network delay
						client := getGitHubClient()
						var err error
						repos, err = searchRepositories(client, searchTerm)
						if err != nil {
							log.Fatalf("Error searching repositories: %v", err)
						}
						loading = false
						showResults = true
						cursorPosition = 0
						startIndex = 0
						focusedElement = 2 // Automatically focus on the results
						s.PostEvent(tcell.NewEventInterrupt(nil)) // Force screen refresh
					}()
				}
			case tcell.KeyUp:
				if focusedElement == 2 && cursorPosition > 0 {
					cursorPosition--
					if cursorPosition < startIndex {
						startIndex--
					}
				}
			case tcell.KeyDown:
				if focusedElement == 2 && cursorPosition < len(repos)-1 {
					cursorPosition++
					if cursorPosition >= startIndex+displayHeight {
						startIndex++
					}
				}
			}
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventInterrupt:
			// Event to forcefully redraw the screen when search completes
		}
	}
}
