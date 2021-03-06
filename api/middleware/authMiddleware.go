package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matthieurobert/amp/api/auth"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing auth middleware")
		_, user, err := auth.Strategy.AuthenticateRequest(r)
		if err != nil {
			fmt.Println(err)
			code := http.StatusUnauthorized
			http.Error(w, http.StatusText(code), code)
			return
		}
		log.Println("User %s Authenticated\n", user.GetUserName())
		next.ServeHTTP(w, r)
	})
}
