package ui

import (
	resources "cwGame/assets/img"
	"image/color"

	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func MainMenu(app fyne.App) fyne.CanvasObject {
	Logo := CreateImage(resources.MainLogo, 150, 150, canvas.ImageFillContain)

	ListPlayers := widget.NewButton("Player list", func() {
		WinListPlayer := app.NewWindow("Player list")
		WinListPlayer.SetFixedSize(true)
		WinListPlayer.Resize(fyne.NewSize(400, 300))
		WinListPlayer.SetContent(ShowListPlayers(app, WinListPlayer))
		WinListPlayer.Show()
	})

	NewGame := widget.NewButton("New game", func() {
		//Create lobby, pick 10 players for game and check who win game
	})
	Leaderboard := widget.NewButton("Leaderboard", func() {
		//check %wins, all played matches
	})

	spacing := canvas.NewRectangle(color.Transparent)
	spacing.SetMinSize(fyne.NewSize(0, 25))

	RightMainBox := container.NewVBox(ListPlayers, spacing, NewGame, spacing, Leaderboard)

	LeftMainBox := container.NewVBox(Logo)

	objs := []fyne.CanvasObject{LeftMainBox, RightMainBox}
	return container.New(NewRemoteLayout(LeftMainBox, RightMainBox), objs...)

}

func CreateImage(path *fyne.StaticResource, width float32, height float32, fillMode canvas.ImageFill) *canvas.Image {
	image := canvas.NewImageFromResource(path)
	image.FillMode = fillMode
	image.SetMinSize(fyne.NewSize(width, height))
	image.Refresh()
	return image
}
