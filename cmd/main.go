package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"test_gyne/internal/layout"
)

func main(){
	a := app.New()
	w := a.NewWindow("Kafka")

	login := layout.NewLogin(w)
	profile := layout.NewProfile(w, login.Get()).Render()
	mainScreen := layout.NewMainScreen(w, profile).Render()

	w.Resize(fyne.NewSize(500,500))
	w.SetContent(login.Render(mainScreen))
	w.ShowAndRun()
}
