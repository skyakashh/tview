package welcome_page

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Welcome() *tview.Box {
	box := tview.NewBox().SetBackgroundColor(tcell.ColorDeepSkyBlue).SetBorder(true)
	return box
}
