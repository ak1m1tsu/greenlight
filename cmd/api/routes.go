package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	moviesRead  = "movies:read"
	moviesWrite = "movies:write"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/movies", app.requirePermission(moviesRead, app.createMovieHandler))
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.requirePermission(moviesWrite, app.listMoviesHandler))
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.requirePermission(moviesRead, app.showMovieHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/movies/:id", app.requirePermission(moviesWrite, app.updateMovieHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.requirePermission(moviesWrite, app.deleteMovieHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/activation", app.createActivationTokenHandler)

	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return app.metrics(app.recoverPanic(app.cors(app.ratelimit(app.authenticate(router)))))
}
