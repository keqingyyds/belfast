package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

func RegisterAdminAuthz(app *iris.Application) {
	handler := handlers.NewAdminAuthzHandler()
	party := app.Party("/api/v1/admin/authz")
	party.Use(middleware.RequirePermissionAny(authz.PermAdminAuthz))
	handlers.RegisterAdminAuthzRoutes(party, handler)
}
