package user

import (
	"cxrus-app/controllers/user"
	"net/http"

	"github.com/gorilla/mux"
)

// InitUserRoutes mapping user controller to path
func InitUserRoutes(r *mux.Router) *mux.Router {

	r.HandleFunc("/api/auth/login", user.Login).Methods(http.MethodPost)
	r.HandleFunc("/api/auth/register", user.Register).Methods(http.MethodPost)
	r.HandleFunc("/api/auth/test", user.Testing).Methods(http.MethodGet)

	return r
}
