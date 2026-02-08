package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

func RegisterDorm3d(app *iris.Application) {
	party := app.Party("/api/v1/dorm3d-apartments")
	party.Use(middleware.RequirePermissionAny(authz.PermDorm3D))
	handler := handlers.NewDorm3dHandler()
	handlers.RegisterDorm3dRoutes(party, handler)
}
