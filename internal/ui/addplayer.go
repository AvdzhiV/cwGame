package ui

import (
	"cwGame/internal/handlers"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func AddNewPlayer(PlayerList *widget.List) fyne.CanvasObject {
	nicknameEntry := widget.NewEntry()
	nicknameEntry.SetPlaceHolder("Player nickname")

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Player name")

	winEntry := widget.NewEntry()
	winEntry.SetPlaceHolder("Player win")

	lossEntry := widget.NewEntry()
	lossEntry.SetPlaceHolder("Player loss")
	createPlayer := widget.NewButton("Create player", func() {
		nickname := nicknameEntry.Text
		name := nameEntry.Text
		win, _ := strconv.Atoi(winEntry.Text)
		loss, _ := strconv.Atoi(lossEntry.Text)

		newPlayer := handlers.Players{
			Name:     name,
			Nickname: nickname,
			Win:      win,
			Loss:     loss,
		}

		handlers.SendPlayerToServer(newPlayer)
		PlayerList.Refresh()
	})

	buttons := container.NewHBox(layout.NewSpacer(), createPlayer, layout.NewSpacer())

	content := container.NewVBox(nicknameEntry, nameEntry, winEntry, lossEntry, buttons)

	return content
}
