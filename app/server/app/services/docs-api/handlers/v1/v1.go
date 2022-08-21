// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"net/http"
	"server/app/services/docs-api/handlers/v1/docgrp"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log *zap.SugaredLogger
	// DB   *sqlx.DB
}

// Routes binds all the version 1 routes.
func Routes(app *gin.Engine, cfg Config) {
	const version = "v1"

	// Register product and sale endpoints.
	dgh := docgrp.Handlers{}
	g := app.Group("/" + version)
	g.Handle(http.MethodGet, "/docs", dgh.Query)
	g.Handle(http.MethodGet, "/docs/:id", dgh.QueryByID)
}
