package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/middleware"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/authz"
)

func RegisterJuustagram(app *iris.Application) {
	handler := handlers.NewJuustagramHandler()
	adminParty := app.Party("/api/v1/juustagram")
	adminParty.Use(middleware.RequirePermissionAny(authz.PermJuustagram))
	handlers.RegisterJuustagramRoutes(adminParty, handler)
	playerParty := app.Party("/api/v1/players/{id:uint}/juustagram")
	playerParty.Use(middleware.RequirePermissionAnyOrSelf(authz.PermPlayers))
	handlers.RegisterJuustagramPlayerRoutes(playerParty, handler)
}
