package actions

import (
	"encoding/json"
	"fmt"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/x/sessions"
	"github.com/jcepedavillamayor/apiexample/models"
	"github.com/rs/cors"
)

type ErrorStruct struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

func errorHandler() buffalo.ErrorHandler {
	return func(status int, err error, c buffalo.Context) error {
		c.Logger().Error(err)
		c.Response().WriteHeader(status)
		msg := fmt.Sprintf("%+v", err.Error())

		return json.NewEncoder(c.Response()).Encode(map[string]interface{}{
			"error":  msg,
			"status": status,
		})

	}

}

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_apiexample_session",
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		app.Use(middleware.PopTransaction(models.DB))
		app.ErrorHandlers[400] = errorHandler()
		app.ErrorHandlers[422] = errorHandler()
		app.ErrorHandlers[500] = errorHandler()

		app.Resource("/frameworks", FrameworksResource{&buffalo.BaseResource{}})
	}

	return app
}
