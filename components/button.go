package components

import "github.com/gdamore/tcell/v2"

func DrawButton(s tcell.Screen, x, y, width int, label string, style tcell.Style, focused bool) {
	if focused {
		style = style.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
	}

	s.SetContent(x, y, '[', nil, style)
	for i := 0; i < width-2; i++ {
		if i < len(label) {
			s.SetContent(x+i+1, y, rune(label[i]), nil, style)
		} else {
			s.SetContent(x+i+1, y, ' ', nil, style)
		}
	}
	s.SetContent(x+width-1, y, ']', nil, style)
	s.Show()
}
