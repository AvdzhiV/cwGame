package ui

import "fyne.io/fyne/v2"

const (
	sideWidth = 800
	TopHeight = 600
)

type RemoteLayout struct {
	left, right fyne.CanvasObject
}

type FormContent struct {
	left, right fyne.CanvasObject
}

func NewRemoteLayout(left, right fyne.CanvasObject) fyne.Layout {
	return &RemoteLayout{left: left, right: right}
}

func (l *RemoteLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	var TopBorderHeight float32 = 20
	var SideBorderWidth float32 = 40
	l.left.Move(fyne.NewPos(SideBorderWidth+20, TopBorderHeight))
	l.left.Resize(fyne.NewSize(300, 300))
	l.right.Move(fyne.NewPos(sideWidth-SideBorderWidth*9, TopBorderHeight+10))
	l.right.Resize(fyne.NewSize(250, 600))
}

func (l *RemoteLayout) MinSize(_ []fyne.CanvasObject) fyne.Size { return fyne.NewSize(800, 600) }
