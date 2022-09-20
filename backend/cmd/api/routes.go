package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/login", app.login)
	router.HandlerFunc(http.MethodPost, "/v1/login-google", app.loginGoogle)
	router.HandlerFunc(http.MethodPost, "/v1/register", app.register)

	router.GET("/v1/washers", app.wrap(secure.ThenFunc(app.getWashers)))
	//router.HandlerFunc(http.MethodGet, "/v1/washers", app.getWashers)
	router.GET("/v1/dryers", app.wrap(secure.ThenFunc(app.getDryers)))
	//router.HandlerFunc(http.MethodGet, "/v1/dryers", app.getDryers)
	router.POST("/v1/create-payment-intent", app.wrap(secure.ThenFunc(app.createPaymentIntent)))
	//router.HandlerFunc(http.MethodPost, "/v1/create-payment-intent", app.createPaymentIntent)
	router.POST("/v1/confirm-payment", app.wrap(secure.ThenFunc(app.confirmPayment)))

	return app.enableCORS(router)
}
