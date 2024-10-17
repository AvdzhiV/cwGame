package ui

import (
	"cwGame/internal/handlers"
	"errors"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowListPlayers(app fyne.App, win fyne.Window) fyne.CanvasObject {
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

	var SelectedPlayer *handlers.Players

	PlayerList.OnSelected = func(id widget.ListItemID) {
		SelectedPlayer = &players[id]
		fmt.Println(SelectedPlayer.Nickname)
	}

	addPlayerButton := widget.NewButton("Create new player", func() {
		winAddNewPlayer := app.NewWindow("Add new player")
		winAddNewPlayer.SetFixedSize(true)
		winAddNewPlayer.Resize(fyne.NewSize(400, 300))
		winAddNewPlayer.SetContent(AddNewPlayer(PlayerList, &players))
		winAddNewPlayer.Show()
	})

	deletePlayerButton := widget.NewButton("Delete player", func() {
		if SelectedPlayer != nil {
			dialog.ShowConfirm("Confirm Delete", "Are you sure you want to delete"+SelectedPlayer.Nickname+"?", func(b bool) {
				if b {
					handlers.DeletePlayerOnServer(SelectedPlayer.Nickname)
					players = handlers.GetPlayersFromServer()

					PlayerList.Length = func() int {
						return len(players)
					}
					PlayerList.UpdateItem = func(id widget.ListItemID, o fyne.CanvasObject) {
						o.(*widget.Label).SetText(players[id].Nickname)
					}
					PlayerList.Refresh()
				}
			}, win)
		} else {
			dialog.ShowError(errors.New("No player selected"), win)
		}

	})

	buttons := container.NewHBox(layout.NewSpacer(), addPlayerButton, deletePlayerButton, layout.NewSpacer())
	content := container.NewBorder(nil, buttons, nil, nil, PlayerList)
	PlayerList.Resize(fyne.NewSize(300, 300))
	return content
}
