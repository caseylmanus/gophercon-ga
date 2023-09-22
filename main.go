package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/caseylmanus/gophercon-ga/queens"
	"github.com/caseylmanus/gophercon-ga/text"
)

func main() {
	//greeting.Solve()
	//queens.Solve()
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	myWindow := myApp.NewWindow("Gophercon 2023 Demo")
	valueLabel := widget.NewTextGrid() //NewLabelWithStyle("", widget.RichTextStyleCodeBlock.Alignment, widget.RichTextStyleCodeBlock.TextStyle)
	start8Queens := widget.NewButton("8 Queens", func() {
		valueLabel.SetText("")
		start := time.Now()
		queens.Solve(8, 1, func(s string) {
			valueLabel.SetText(s)
		})
		valueLabel.SetText(valueLabel.Text() + fmt.Sprintln("Completed in: ", time.Since(start)))
		//valueLabel.Refresh()
	})
	start16Queens := widget.NewButton("16 Queens (4) species", func() {
		valueLabel.SetText("")
		start := time.Now()
		queens.Solve(16, 4, func(s string) {
			valueLabel.SetText(valueLabel.Text() + s)
		})
		valueLabel.SetText(valueLabel.Text() + fmt.Sprintln("Completed in: ", time.Since(start)))
		//valueLabel.Refresh()
	})
	startGreeting := widget.NewButton("Hello Gophers!", func() {
		valueLabel.SetText("")
		start := time.Now()
		target := "Hello Gophercon 2023, Welcome to San Diego!"
		text.Solve(target, func(s string) {
			valueLabel.SetText(valueLabel.Text() + s)
		})
		valueLabel.SetText(valueLabel.Text() + fmt.Sprintln("Completed in:", time.Since(start)))
		//valueLabel.Refresh()
	})
	buttons := container.New(layout.NewHBoxLayout(), startGreeting, start8Queens, start16Queens)

	content := container.New(layout.NewVBoxLayout(), buttons, valueLabel)
	myWindow.Resize(fyne.NewSize(1028, 764))
	//myWindow.SetFullScreen(true)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
