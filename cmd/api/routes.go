package api

import (
	"github.com/gorilla/mux"
	"github.com/juju/ratelimit"
	"github.com/mxnyawi/doesItSponsor/internal/handler"
)

func SetRoutes(router *mux.Router, limiter *ratelimit.Bucket, handler *handler.Handler) {
	router.HandleFunc("/organisation/{organisation_name}", limitMiddleware(limiter, handler.GetOrganisation)).Methods("GET")
	router.HandleFunc("/county/{county}", limitMiddleware(limiter, handler.GetCounty)).Methods("GET")
	router.HandleFunc("/type/{type}", limitMiddleware(limiter, handler.GetType)).Methods("GET")
	router.HandleFunc("/route/{route}", limitMiddleware(limiter, handler.GetRoute)).Methods("GET")
	router.HandleFunc("/city/{city}", limitMiddleware(limiter, handler.GetCity)).Methods("GET")
}
