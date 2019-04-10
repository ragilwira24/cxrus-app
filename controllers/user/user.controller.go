package user

import (
	"cxrus-app/service/impl"
	"cxrus-app/service/models"
	"cxrus-app/util/handler"
	"encoding/json"
	"net/http"
)

// Register Controller
var Register = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		handler.Message(false, "Invalid Request")
		return
	}

	response := impl.Register(user)
	handler.Response(w, response)

}

// Login Controller
var Login = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		handler.Message(false, "Invalid Request")
		return
	}

	response := impl.Login(user.Email, user.Password)
	handler.Response(w, response)
}

// Testing controller
var Testing = func(w http.ResponseWriter, r *http.Request) {
	response := handler.Message(true, "TESTING")
	handler.Response(w, response)
}
