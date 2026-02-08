package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

func RegisterPlayers(app *iris.Application) {
	party := app.Party("/api/v1/players")
	party.Use(middleware.RequirePermissionAnyOrSelf(authz.PermPlayers))
	handler := handlers.NewPlayerHandler()
	handlers.RegisterPlayerRoutes(party, handler)
}
