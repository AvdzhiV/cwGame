package resources

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed dota2-logo-official.png
var Logo []byte

var MainLogo = &fyne.StaticResource{
	StaticName:    "dota2-logo-official.png",
	StaticContent: Logo,
}
