package components

import "github.com/gdamore/tcell/v2"

// DrawText is a helper function to draw text on the screen at a specific position.
func DrawText(s tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range text {
		s.SetContent(x+i, y, r, nil, style)
	}
	s.Show()
}
