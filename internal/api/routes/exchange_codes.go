package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

func RegisterExchangeCodes(app *iris.Application) {
	party := app.Party("/api/v1/exchange-codes")
	party.Use(middleware.RequirePermissionAny(authz.PermExchangeCodes))
	handler := handlers.NewExchangeCodeHandler()
	handlers.RegisterExchangeCodeRoutes(party, handler)
}
