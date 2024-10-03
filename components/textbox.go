package components

import "github.com/gdamore/tcell/v2"

func DrawTextBox(s tcell.Screen, x, y, width int, text string, style tcell.Style) {
	for i := 0; i < width; i++ {
		s.SetContent(x+i, y, tcell.RuneHLine, nil, style)
	}
	s.SetContent(x, y, tcell.RuneULCorner, nil, style)
	s.SetContent(x+width-1, y, tcell.RuneURCorner, nil, style)

	for i, r := range text {
		if i >= width-2 {
			break
		}
		s.SetContent(x+1+i, y, r, nil, style)
	}
	s.Show()
}
