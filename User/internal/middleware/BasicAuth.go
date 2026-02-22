package middleware

import (
	"User/internal/utils"
	"crypto/subtle"
	"net/http"
)

func BasicAuth(username, password string, realm string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
				utils.AuthErrorResponse(w)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
