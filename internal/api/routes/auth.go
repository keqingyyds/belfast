package routes

import (
	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/go-webauthn/webauthn/protocol"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/handlers"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/auth"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/config"
	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/logger"
)

func RegisterAuth(app *iris.Application, cfg *config.Config) *auth.Manager {
	if cfg == nil {
		cfg = &config.Config{}
	}
	manager, err := auth.NewManager(cfg.Auth)
	if err != nil {
		logger.LogEvent("API", "Auth", "webauthn disabled: "+err.Error(), logger.LOG_LEVEL_WARN)
		manager = &auth.Manager{Config: auth.NormalizeConfig(cfg.Auth), Limiter: auth.NewRateLimiter(), Selection: protocol.AuthenticatorSelection{UserVerification: protocol.VerificationPreferred}}
	}
	handler := handlers.NewAuthHandler(manager)
	party := app.Party("/api/v1/auth")
	handlers.RegisterAuthRoutes(party, handler)
	return manager
}
