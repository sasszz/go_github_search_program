package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/google/go-github/v65/github"
)

// DrawGitHubRepositories displays the list of GitHub repositories using tcell.
func DrawGitHubRepositories(s tcell.Screen, x, y, width, height int, repos []*github.Repository, startIndex, selectedIndex int, style tcell.Style) {
	for i := 0; i < height; i++ {
		itemIndex := startIndex + i
		if itemIndex >= len(repos) {
			break
		}
		repo := repos[itemIndex]
		// Format the repository information
		item := fmt.Sprintf("Name: %s | Owner: %s | Stars: %d", repo.GetName(), repo.Owner.GetLogin(), repo.GetStargazersCount())
		if itemIndex == selectedIndex {
			style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
		} else {
			style = style.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)
		}
		// Render each repository item
		for j, r := range []rune(item) {
			if j < width {
				s.SetContent(x+j, y+i, r, nil, style)
			}
		}
	}
	s.Show()
}
