package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/config"
)

func RegisterServer(app *iris.Application, cfg *config.Config) {
	party := app.Party("/api/v1/server")
	handler := &handlers.ServerHandler{Config: cfg}
	handlers.RegisterServerRoutes(party, handler)
}
