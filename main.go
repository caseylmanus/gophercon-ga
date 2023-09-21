package main

import (
	"fmt"
	"time"

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
	startQueens := widget.NewButton("8 Queens!", func() {
		valueLabel.SetText("")
		start := time.Now()
		queens.Solve(func(s string) {
			valueLabel.SetText(valueLabel.Text + s)
		})
		valueLabel.SetText(valueLabel.Text + fmt.Sprintln("Completed in: ", time.Since(start)))
	})
	startGreeting := widget.NewButton("Hello Gophers!", func() {
		valueLabel.SetText("")
		start := time.Now()
		greeting.Solve(func(s string) {
			valueLabel.SetText(valueLabel.Text + s)
		})
		valueLabel.SetText(valueLabel.Text + fmt.Sprintln("Completed in:", time.Since(start)))
	})
	buttons := container.New(layout.NewHBoxLayout(), startGreeting, startQueens)
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
