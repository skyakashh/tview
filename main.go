package main

import "github.com/rivo/tview"
import Welcome

func main() {

	if err := tview.NewApplication().SetRoot(Welcome(), true).Run(); err != nil {
		panic(err)
	}
}
