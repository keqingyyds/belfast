package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/auth"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

func RegisterAdminUsers(app *iris.Application, manager *auth.Manager) {
	handler := handlers.NewAdminUserHandler(manager)
	party := app.Party("/api/v1/admin/users")
	party.Use(middleware.RequirePermissionAny(authz.PermAdminUsers))
	handlers.RegisterAdminUserRoutes(party, handler)
}
