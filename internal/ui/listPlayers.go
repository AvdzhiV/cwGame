package ui

import (
	"cwGame/internal/handlers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowListPlayers(app fyne.App) fyne.CanvasObject {
	players := handlers.GetPlayersFromServer()

	playerList := widget.NewList(
		func() int {
			return len(players)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Player")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(players[i].Nickname)
		},
	)

	addPlayerButton := widget.NewButton("Add player", func() {
		newPlayer := handlers.Players{
			Nickname: "Player",
		}

		handlers.SendPlayerToServer(newPlayer)

		players = handlers.GetPlayersFromServer()
		playerList.Refresh()
	})
	buttons := container.NewHBox(addPlayerButton)
	content := container.NewVBox(playerList, buttons)
	return content
}
