package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"cwGame/internal/ui"
)

const (
	appId   = "com.dota.cw"
	appName = "cwGame"
)

var (
	AppDota fyne.App
	winDota fyne.Window
)

func main() {
	AppDota = app.NewWithID(appId)
	winDota = AppDota.NewWindow(appName)
	winDota.SetFixedSize(true)
	winDota.CenterOnScreen()
	winDota.Resize(fyne.NewSize(800, 600))
	winDota.SetContent(ui.MainMenu(AppDota))
	winDota.ShowAndRun()
}
