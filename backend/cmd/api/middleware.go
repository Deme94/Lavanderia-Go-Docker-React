package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (app *application) enableCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		handler.ServeHTTP(w, r)
	})
}

func (app *application) checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			// could set an anonoymous user
			app.errorJSON(w, errors.New("invalid authorization header"))
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			app.errorJSON(w, errors.New("invalid authorization header"))
			return
		}

		if headerParts[0] != "Bearer" {
			app.errorJSON(w, errors.New("unauthorized - no bearer"))
			return
		}

		token := headerParts[1]
		claims, err := jwt.HMACCheck([]byte(token), []byte(app.config.jwt.secret))
		if err != nil {
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		if !claims.Valid(time.Now()) {
			app.errorJSON(w, errors.New("unauthorized - token expired"), http.StatusForbidden)
			return
		}

		if !claims.AcceptAudience(domain) {
			app.errorJSON(w, errors.New("unauthorized - invalid audience"), http.StatusForbidden)
			return
		}

		if claims.Issuer != domain {
			app.errorJSON(w, errors.New("unauthorized - invalid issuer"), http.StatusForbidden)
			return
		}

		app.logger.Println("Valid user:", claims.Subject)

		next.ServeHTTP(w, r)
	})
}
