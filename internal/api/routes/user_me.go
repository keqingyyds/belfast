package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/config"
)

func RegisterMe(app *iris.Application, cfg *config.Config) {
	handler := handlers.NewMeHandler()
	party := app.Party("/api/v1/me")
	handlers.RegisterMeRoutes(party, handler)
}
