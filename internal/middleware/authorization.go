package middleware

import (
	"errors"
	"net/http"

	"github.com/Abhishekkapoor98/goapi/api"
	"github.com/Abhishekkapoor98/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizationError = errors.New("Invalid Username or Token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		// Simple Authorization Logic
		if username == "" || token == "" {
			log.Error(UnAuthorizationError)
			api.RequestErrorHandler(w, UnAuthorizationError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDeatils
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizationError)
			api.RequestErrorHandler(w, UnAuthorizationError)
			return
		}

		// Proceed to next handler if authorized
		next.ServeHTTP(w, r)
	})
}
