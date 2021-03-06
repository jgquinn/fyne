package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func makeButtonTab() fyne.Widget {
	return widget.NewVBox(
		widget.NewLabel("Text label"),
		widget.NewButton("Text button", func() { fmt.Println("tapped text button") }),
		widget.NewButtonWithIcon("With icon", theme.ConfirmIcon(), func() { fmt.Println("tapped icon button") }),
	)
}

func makeInputTab() fyne.Widget {
	entry := widget.NewEntry()
	entry.SetText("Entry")

	return widget.NewVBox(
		entry,
		widget.NewCheck("Check", func(on bool) { fmt.Println("checked", on) }),
		widget.NewRadio([]string{"Item 1", "Item 2"}, func(s string) { fmt.Println("selected", s) }),
	)
}

// Widget shows a window containing widget demos
func Widget(app fyne.App) {
	w := app.NewWindow("Widgets")

	w.SetContent(widget.NewTabContainer(
		widget.NewTabItem("Buttons", makeButtonTab()),
		widget.NewTabItem("Input", makeInputTab()),
		widget.NewTabItem("Group",
			widget.NewGroup("Grouped", widget.NewLabel("Grouped content"))),
	))

	w.Show()
}
