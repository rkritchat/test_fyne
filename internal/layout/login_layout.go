package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Login interface {
	Render(next *fyne.Container) *fyne.Container
	Get() func() (*binding.String, *binding.String)
}

type login struct {
	w        fyne.Window
	Username binding.String
	Password binding.String
}

func NewLogin(w fyne.Window) Login {
	return &login{
		w:        w,
		Username: binding.NewString(),
		Password: binding.NewString(),
	}
}

func (l login) Render(next *fyne.Container) *fyne.Container {
	l.Username.Set("a")
	l.Password.Set("b")
	usernameBox := widget.NewEntryWithData(l.Username)
	usernameBox.Resize(fyne.NewSize(250, 20))
	pwdBox := widget.NewEntryWithData(l.Password)
	pwdBox.Resize(fyne.NewSize(250, 20))
	return container.NewVBox(
		container.NewHBox(widget.NewLabel("Username:"), usernameBox),
		container.NewHBox(widget.NewLabel("Password:"), pwdBox),
		widget.NewButton("Login", func() {
			l.handler(next)
		}),
	)
}

func (l login) handler(next *fyne.Container) {
	usernameStr, _ := l.Username.Get()
	pwdStr, _ := l.Password.Get()
	if usernameStr == "a" && pwdStr == "b" {
		l.next(next)
	}
}

func (l login) next(next *fyne.Container) {
	l.w.SetContent(next)
}

func (l login) Get() func() (*binding.String, *binding.String) {
	return func() (*binding.String, *binding.String) { return &l.Username, &l.Password }
}
