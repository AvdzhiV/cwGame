package ui

import (
	resources "cwGame/assets/img"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func MainMenu() fyne.CanvasObject {
	Logo := CreateImage(resources.Logo, 100, 100, canvas.ImageFillContain)

	//RightMainBox := container.NewVBox("Список игроков", "Новая игра", "Таблица лидеров")

	LeftMainBox := container.NewVBox(Logo)

	return container.NewStack(LeftMainBox)

}

func CreateImage(path *fyne.StaticResource, width float32, height float32, fillMode canvas.ImageFill) *canvas.Image {
	image := canvas.NewImageFromResource(path)
	image.FillMode = fillMode
	image.SetMinSize(fyne.NewSize(width, height))
	image.Refresh()
	return image
}
