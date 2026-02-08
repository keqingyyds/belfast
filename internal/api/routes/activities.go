package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

func RegisterActivities(app *iris.Application) {
	party := app.Party("/api/v1/activities")
	party.Use(middleware.RequirePermissionAny(authz.PermActivities))
	handler := handlers.NewActivityHandler()
	handlers.RegisterActivityRoutes(party, handler)
}
