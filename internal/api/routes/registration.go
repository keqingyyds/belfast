package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/config"
)

func RegisterRegistration(app *iris.Application, cfg *config.Config) {
	handler := handlers.NewRegistrationHandler(cfg)
	party := app.Party("/api/v1/registration")
	handlers.RegisterRegistrationRoutes(party, handler)
}
