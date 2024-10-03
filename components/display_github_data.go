package components

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/google/go-github/v65/github"
)

// DrawGitHubRepositories displays the list of GitHub repositories using tcell with truncated text and ellipsis to fit uniform columns.
func DrawGitHubRepositories(s tcell.Screen, x, y, width, height int, repos []*github.Repository, startIndex, selectedIndex int, style tcell.Style) {
	// Column widths for each field (adjusted to allow for ellipsis)
	nameWidth := 14  // Allows 12 characters plus ellipsis (if needed)
	ownerWidth := 12 // Allows 10 characters plus ellipsis (if needed)
	starsWidth := 6
	issuesWidth := 10
	prsWidth := 6

	// Draw table header
	header := fmt.Sprintf("| %-*s | %-*s | %-*s | %-*s | %-*s |",
		nameWidth, "NAME",
		ownerWidth, "OWNER",
		starsWidth, "STARS",
		issuesWidth, "# of Issues",
		prsWidth, "# of PRs")
	drawText(s, x, y, header, style)
	drawLine(s, x, y+1, width, "=", style)

	// Display repositories in the table format
	for i := 0; i < height && startIndex+i < len(repos); i++ {
		repo := repos[startIndex+i]
		name := truncateStringWithEllipsis(repo.GetName(), nameWidth)
		owner := truncateStringWithEllipsis(repo.Owner.GetLogin(), ownerWidth)
		stars := fmt.Sprintf("%d", repo.GetStargazersCount())
		issues := fmt.Sprintf("%d", repo.GetOpenIssuesCount())
		prs := "N/A" // Replace with the actual PR count if available

		// Format row text
		rowText := fmt.Sprintf("| %-*s | %-*s | %-*s | %-*s | %-*s |",
			nameWidth, name,
			ownerWidth, owner,
			starsWidth, stars,
			issuesWidth, issues,
			prsWidth, prs)

		// Highlight the selected row
		if startIndex+i == selectedIndex {
			drawText(s, x, y+2+i, ">"+rowText[1:], style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite))
		} else {
			drawText(s, x, y+2+i, rowText, style)
		}
	}

	// Draw bottom of the table
	drawLine(s, x, y+height+2, width, "-", style)

	s.Show()
}

// truncateStringWithEllipsis ensures the text fits within a given width, adding ellipsis if necessary.
func truncateStringWithEllipsis(text string, maxLen int) string {
	if len(text) > maxLen {
		return text[:maxLen-3] + "..." // Truncate and add ellipsis if too long
	}
	return text
}

// Helper function to draw text on the screen
func drawText(s tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, style)
	}
}

// Helper function to draw a horizontal line of a specific character
func drawLine(s tcell.Screen, x, y, length int, char string, style tcell.Style) {
	for i := 0; i < length; i++ {
		s.SetContent(x+i, y, rune(char[0]), nil, style)
	}
}
