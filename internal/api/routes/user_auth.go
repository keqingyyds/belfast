package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/config"
)

func RegisterUserAuth(app *iris.Application, cfg *config.Config) {
	handler := handlers.NewUserAuthHandler(cfg)
	party := app.Party("/api/v1/user/auth")
	handlers.RegisterUserAuthRoutes(party, handler)
}
