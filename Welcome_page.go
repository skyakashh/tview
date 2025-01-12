package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func Welcome() *tview.Box {
	Box := tview.NewBox().SetBackgroundColor(tcell.ColorDeepSkyBlue).SetBorder(true)
	return Box
}
