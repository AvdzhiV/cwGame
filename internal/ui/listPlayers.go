package ui

import (
	"cwGame/internal/handlers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowListPlayers(app fyne.App) fyne.CanvasObject {
	players := handlers.GetPlayersFromServer()

	PlayerList := widget.NewList(
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

	addPlayerButton := widget.NewButton("Create new player", func() {
		winAddNewPlayer := app.NewWindow("Add new player")
		winAddNewPlayer.SetFixedSize(true)
		winAddNewPlayer.Resize(fyne.NewSize(400, 300))
		winAddNewPlayer.SetContent(AddNewPlayer(PlayerList))
		winAddNewPlayer.Show()
	})

	buttons := container.NewHBox(layout.NewSpacer(), addPlayerButton, layout.NewSpacer())
	content := container.NewBorder(nil, buttons, nil, nil, PlayerList)
	PlayerList.Resize(fyne.NewSize(300, 300))
	return content
}

/* Кнопка отправки на сервер нового пользователя
newPlayer := handlers.Players{
	Nickname: "Player",
}

handlers.SendPlayerToServer(newPlayer)

players = handlers.GetPlayersFromServer()
playerList.Refresh()
*/
