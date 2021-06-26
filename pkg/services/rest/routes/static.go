package routes

import (
	"net/http"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/inits"
	"github.com/gigamono/gigamono/pkg/services/rest/middleware"
	"github.com/gin-gonic/gin"
)

// SetLocalStaticRoutes sets local static file routes depending on config.
func SetLocalStaticRoutes(server *gin.Engine, app *inits.App) {
	// Local static folder to serve project files.
	if app.Config.Filestore.Project.Kind == configs.Local {
		// TODO: Permission middleware.
		// Authenticate and add session user.
		staticRoute := server.Group("/"+app.Config.Filestore.Project.Paths.Public, middleware.AuthenticateCreateUser(app))
		staticRoute.StaticFS("/", http.Dir(app.Config.Filestore.Project.Paths.Private))
	}

	// Local static folder to serve extension files.
	if app.Config.Filestore.Extension.Kind == configs.Local {
		// TODO: Permission middleware.
		// Authenticate and add session user.
		staticRoute := server.Group("/"+app.Config.Filestore.Extension.Paths.Public, middleware.AuthenticateCreateUser(app))
		staticRoute.StaticFS("/", http.Dir(app.Config.Filestore.Extension.Paths.Private))
	}

	// Local static folder to serve image files.
	if app.Config.Filestore.Image.Kind == configs.Local {
		// TODO: Permission middleware.
		// Authenticate and add session user.
		staticRoute := server.Group("/"+app.Config.Filestore.Image.Paths.Public, middleware.AuthenticateCreateUser(app))
		staticRoute.StaticFS("/", http.Dir(app.Config.Filestore.Image.Paths.Private))
	}
}
