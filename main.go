package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/google/go-github/v65/github"
)

func drawText(s tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range []rune(text) {
		s.SetContent(x+i, y, r, nil, style)
	}
}

func displayRepositories(s tcell.Screen, repos []*github.Repository, style tcell.Style) {
	s.Clear()
	y := 0
	for _, repo := range repos {
		text := fmt.Sprintf("Name: %s | Owner: %s | Stars: %d | Description: %s",
			repo.GetName(), repo.Owner.GetLogin(), repo.GetStargazersCount(), repo.GetDescription())
		drawText(s, 0, y, text, style)
		y++
		if y >= 20 {
			break
		}
	}
	s.Show()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gx <search_term>")
		return
	}

	searchTerm := os.Args[1]
	client := getGitHubClient()
	repos, err := searchRepositories(client, searchTerm)
	if err != nil {
		log.Fatalf("Error searching repositories: %v", err)
	}

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.Clear()

	displayRepositories(s, repos, defStyle)

	// Event loop
	for {
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				s.Fini()
				return
			}
		}
	}
}
