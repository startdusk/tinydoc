// Package handlers manages the different versions of the API.
package handlers

import (
	"net/http"
	"os"
	v1 "server/app/services/docs-api/handlers/v1"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	// Construct the  which holds all routes as well as common Middleware.
	var app *gin.Engine

	// Do we need CORS?
	if opts.corsOrigin != "" {
		app = gin.New()
		// app = web.NewApp(
		// 	cfg.Shutdown,
		// 	mid.Logger(cfg.Log),
		// 	mid.Errors(cfg.Log),
		// 	mid.Metrics(),
		// 	mid.Cors(opts.corsOrigin),
		// 	mid.Panics(),
		// )

		// Accept CORS 'OPTIONS' preflight requests if config has been provided.
		// Don't forget to apply the CORS middleware to the routes that need it.
		// Example Config: `conf:"default:https://MY_DOMAIN.COM"`
		// h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		// 	return nil
		// }
		// app.Handle(http.MethodOptions, "", "/*", h, mid.Cors(opts.corsOrigin))
	}

	if app == nil {
		app = gin.New()
		// 用日志
		// app = web.NewApp(
		// 	cfg.Shutdown,
		// 	mid.Logger(cfg.Log),
		// 	mid.Errors(cfg.Log),
		// 	mid.Panics(),
		// )
	}

	// Load the v1 routes.
	v1.Routes(app, v1.Config{
		Log: cfg.Log,
	})

	return app
}
