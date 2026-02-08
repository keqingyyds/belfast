package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

const apiBasePath = "/api/v1"

func RegisterGameData(app *iris.Application) {
	gameDataHandler := handlers.NewGameDataHandler()
	party := app.Party(apiBasePath)
	party.Use(middleware.RequirePermissionAny(authz.PermGameData))
	handlers.RegisterGameDataRoutes(party, gameDataHandler)
}
