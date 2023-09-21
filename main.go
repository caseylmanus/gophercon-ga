package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/caseylmanus/gophercon-ga/greeting"
	"github.com/caseylmanus/gophercon-ga/queens"
)

func main() {
	//greeting.Solve()
	//queens.Solve()
	myApp := app.New()
	myWindow := myApp.NewWindow("Gophercon 2023 Demo")
	valueLabel := widget.NewLabel("")
	startQueens := widget.NewButton("Queens!", func() {
		valueLabel.SetText("")
		queens.Solve(func(s string) {
			valueLabel.SetText(valueLabel.Text + s)
		})
	})
	startGreeting := widget.NewButton("Greeting!", func() {
		valueLabel.SetText("")
		greeting.Solve(func(s string) {
			valueLabel.SetText(valueLabel.Text + s)
		})
	})
	buttons := container.New(layout.NewHBoxLayout(), startQueens, startGreeting)
	//generationLabel := widget.NewLabel("Generation:")
	//fitnessLabel := widget.NewLabel("Max Fitness:")

	content := container.New(layout.NewVBoxLayout(), buttons, valueLabel) //, generationLabel, fitnessLabel, valueLabel)

	//content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
	//	log.Println("tapped home")
	//})
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(1200, 900))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()

}
