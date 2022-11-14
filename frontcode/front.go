package frontcode

import (
	"goappex"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.Div().Class("image-title").Body(
			app.Img().
				Alt("butterfly").
				Src("/web/logo2.png").
				Width(400).Height(300),
		),
		app.Div().Body(
			app.H1().Class("hello-title").Text("Hello beutiful World! 2"),
		))
}

func init() {
	goappex.Mainfront = mainfront
}

func mainfront() {

	hF := &hello{}
	app.Route(goappex.HelloPath, hF)
	app.Route(goappex.WebrtcExPath, newWebrtcEx())

	app.RunWhenOnBrowser()

}
