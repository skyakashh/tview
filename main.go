package main

import (
	"github.com/rivo/tview"
	"github.com/username/Tview/pages/welcome_page"
)

func main() {

	if err := tview.NewApplication().SetRoot(welcome_page.Welcome(), true).Run(); err != nil {
		panic(err)
	}
}
