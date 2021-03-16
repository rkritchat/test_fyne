package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
const(
	_actionProfile = "PF"
)
type MainScreen interface {
	Render() *fyne.Container
}

type mainScreen struct {
	w       fyne.Window
	profile *fyne.Container
}

func NewMainScreen(w fyne.Window, profile *fyne.Container) MainScreen {
	return &mainScreen{
		w:       w,
		profile: profile,
	}
}

func (m mainScreen) Render() *fyne.Container {
	return container.NewHBox(
		widget.NewButton("Update profile", func() {
			m.next(_actionProfile)
		}),
	)
}

func (m mainScreen) next(ac string){
	switch ac {
	case _actionProfile:
		m.w.SetContent(m.profile)
	}
}
