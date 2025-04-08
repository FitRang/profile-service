package routes

import (
	"net/http"

	"github.com/FitRang/profile-service/handlers"
	"github.com/gin-gonic/gin"
)
type Route struct {
	Name		string
	Method		string
	Pattern		string
	Protected	bool 
	HandlerFunc gin.HandlerFunc
}
type Routes []Route
// Add a health handler later
func NewRoutes(profileHandler *handlers.ProfileHandler) Routes {
	return Routes {
		Route{
			"Create Profile",
			http.MethodPost,
			"/profile",
			false,
			profileHandler.createProfile,
		},	
		Route{
			"Get Profile",
			http.MethodGet,
			"/profile",
			false,
			profileHandler.getProfile,
		},	

	}
}
