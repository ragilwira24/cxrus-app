package auth

import (
	"context"
	"cxrus-app/util/auth"
	authModel "cxrus-app/util/auth/models"
	util "cxrus-app/util/handler"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// JwtAuthentication function
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := auth.AuthorizedPath()
		requestPath := r.URL.Path

		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}
		response := make(map[string]interface{})
		tokenHeader := r.Header.Get("X-Access-Token")

		if tokenHeader == "" {
			response = util.Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Response(w, response)
		}

		tk := &authModel.Token{}
		token, err := jwt.ParseWithClaims(tokenHeader, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token.password")), nil
		})

		if err != nil {
			response = util.Message(false, "Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Response(w, response)
			return
		}

		if !token.Valid {
			response = util.Message(false, "Token is not valid")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			util.Response(w, response)
			return
		}

		type key string
		const keyParam key = "user"

		f := fmt.Sprintf("User %d", tk.UserID) //Useful for monitoring
		fmt.Println(f)
		ctx := context.WithValue(r.Context(), keyParam, tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
