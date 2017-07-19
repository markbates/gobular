package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"

	"github.com/gobuffalo/envy"

	"github.com/gobuffalo/buffalo/middleware/i18n"
	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App
var T *i18n.Translator

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env:         ENV,
			SessionName: "_gobular_session",
		})

		app.Use(middleware.AddContentType("text/html"))
		// Automatically save the session if the underlying
		// Handler does not return an error.
		app.Use(middleware.SessionSaver)
		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("year", time.Now().Year())
				return next(c)
			}
		})

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		// Setup and use translations:
		var err error
		if T, err = i18n.New(packr.NewBox("../locales"), "en-US"); err != nil {
			app.Stop(err)
		}
		app.Use(T.Middleware())

		app.GET("/", NewChecker)
		app.POST("/check", RunChecker)
		app.GET("/check", RunChecker)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
	}

	return app
}
