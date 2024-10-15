package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const (
	appId   = "com.dota.cw"
	appName = "cwGame"
)

var (
	appDota fyne.App
	winDota fyne.Window
)

func main() {
	appDota = app.NewWithID(appId)
	winDota = appDota.NewWindow(appName)
	winDota.SetFixedSize(true)
	winDota.CenterOnScreen()
	winDota.Resize(fyne.NewSize(800, 600))
	winDota.SetContent(ui.MainMenu)
	winDota.ShowAndRun()
}
