package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Profile interface {
	Render() *fyne.Container
}

type profile struct {
	w        fyne.Window
	username *binding.String
	password *binding.String
}

func NewProfile(w fyne.Window, f func() (*binding.String, *binding.String)) Profile {
	u, p := f()
	return &profile{
		w:        w,
		username: u,
		password: p,
	}
}

func (p profile) Render() *fyne.Container {
	//user, _ := (*p.username).Get()
	//pwd, _ := (*p.password).Get()
	return container.NewVBox(
		widget.NewLabel("a"),
		widget.NewLabel("bbbbb"),
	)
}
