package routes

import (
	"time"

	"https://gh.xmly.dev/github.com/kataras/iris/v12"

	"https://gh.xmly.dev/github.com/ggmolly/belfast/internal/api/response"
)

type healthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

// Health godoc
// @Summary     Health check
// @Tags        Health
// @Produce     json
// @Success     200  {object}  healthResponse
// @Router      /health [get]
func Health(ctx iris.Context) {
	payload := healthResponse{
		Status: "ok",
		Time:   time.Now().UTC().Format(time.RFC3339),
	}
	_ = ctx.JSON(response.Success(payload))
}

func Register(app *iris.Application) {
	app.Get("/health", Health)
}
