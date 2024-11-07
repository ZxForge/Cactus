package middleware

import (
	"fmt"
	"net/http"
)

func AuthSession() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Cookies())
			next.ServeHTTP(rw, r)
		})
	}
}
