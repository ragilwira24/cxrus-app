package routes

import (
	"cxrus-app/controllers/auth"
	"cxrus-app/routes/user"

	"github.com/gorilla/mux"
)

// InitBaseRoutes = Init all Base Routes
func InitBaseRoutes() *mux.Router {

	r := mux.NewRouter()
	r = user.InitUserRoutes(r)

	r.Use(auth.JwtAuthentication)
	return r

}
