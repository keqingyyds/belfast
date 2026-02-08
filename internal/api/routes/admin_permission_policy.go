package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

func RegisterAdminPermissionPolicy(app *iris.Application) {
	handler := handlers.NewAdminPermissionPolicyHandler()
	party := app.Party("/api/v1/admin/permission-policy")
	party.Use(middleware.RequirePermissionAny(authz.PermAdminPermission))
	handlers.RegisterAdminPermissionPolicyRoutes(party, handler)
}
