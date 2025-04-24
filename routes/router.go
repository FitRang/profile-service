package routes

import (
	"net/http"

	"github.com/FitRang/profile-service/handlers"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Protected   bool
	HandlerFunc gin.HandlerFunc
}
type Routes []Route

// Add a health handler later
func NewRoutes(profileHandler *handlers.ProfileHandler) Routes {
	return Routes{
		Route{
			"Create Profile",
			http.MethodPost,
			"/profile",
			false,
			profileHandler.CreateProfileHandler,
		},
		Route{
			"Get Profile",
			http.MethodGet,
			"/profile/:id",
			false,
			profileHandler.GetProfileHandler,
		},
		Route{
			"Patch Profile",
			http.MethodPatch,
			"/profile",
			false,
			profileHandler.UpdateProfileHandler,
		},
	}
}

func AttachRoutes(server *gin.Engine, routes Routes) {
	for _, route := range routes {
		if route.Protected {
			server.Handle(route.Method, route.Pattern, route.HandlerFunc)
		} else {
			server.Handle(route.Method, route.Pattern, route.HandlerFunc)
		}
	}
}
