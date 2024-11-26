package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s, Headers: %v", r.Method, r.URL, r.Header)
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(jwtKey []byte) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {

				http.Error(w, "No autenticado", http.StatusUnauthorized)
				return
			}
			tokenStr := cookie.Value
			claims := &jwt.MapClaims{}
			_, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil {
				http.Error(w, "Token inválido o expirado", http.StatusUnauthorized)
				return
			}
			email := (*claims)["email"].(string)
			userId := (*claims)["UserID"].(float64)

			idUser := uint(userId)
			ctx := context.WithValue(r.Context(), "email", email)
			ctx = context.WithValue(ctx, "UserID", idUser)

			next.ServeHTTP(w, r.WithContext(ctx))
		})

	}
}
