package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("こんにちわーるど")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
