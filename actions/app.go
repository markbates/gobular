package actions

import (
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/csrf"
	"github.com/markbates/gobular/models"

	"github.com/gobuffalo/envy"

	"github.com/gobuffalo/packr"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:         ENV,
			SessionName: "_gobular_session",
		})

		app.Use(csrf.New)
		app.Use(middleware.PopTransaction(models.DB))
		app.Use(middleware.AddContentType("text/html"))
		app.Use(func(next buffalo.Handler) buffalo.Handler {
			return func(c buffalo.Context) error {
				c.Set("year", time.Now().Year())
				return next(c)
			}
		})

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		app.GET("/", NewChecker)
		app.POST("/x", RunChecker)
		app.GET("/x/{expression_id}", ReRunChecker)

		app.ServeFiles("/assets", packr.NewBox("../public/assets"))
	}

	return app
}
