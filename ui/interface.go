package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var topWindow fyne.Window

func StartUI() {
	a := app.NewWithID("party.tracker")
	w := a.NewWindow("Party Tracker")
	topWindow = w

	w.SetMaster()

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.Resize(fyne.NewSize(640, 460))

	w.ShowAndRun()
}
