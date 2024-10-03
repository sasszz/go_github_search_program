package components

import "github.com/gdamore/tcell/v2"

func DrawScrollableList(s tcell.Screen, x, y, width, height int, items []string, startIndex, selectedIndex int, style tcell.Style) {
	for i := 0; i < height; i++ {
		itemIndex := startIndex + i
		if itemIndex >= len(items) {
			break
		}
		item := items[itemIndex]
		if itemIndex == selectedIndex {
			style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
		} else {
			style = style.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)
		}
		for j, r := range []rune(item) {
			if j < width {
				s.SetContent(x+j, y+i, r, nil, style)
			}
		}
	}
	s.Show()
}
