package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *server) routes() {
	api := s.router.PathPrefix("/v1/spot").Subrouter()
	api.HandleFunc("/ping", s.PingHandler).Methods(http.MethodGet)
	api.HandleFunc("/version", s.VersionHandler).Methods(http.MethodGet)
	api.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	// elastigroup handlers
	api.HandleFunc("/{account}/elastigroups", s.ElastigroupsListHandler).Methods(http.MethodGet)
	api.HandleFunc("/{account}/elastigroups", s.ElastigroupCreateHandler).Methods(http.MethodPost)
	api.HandleFunc("/{account}/elastigroups/{elastigroup}", s.ElastigroupUpdateHandler).Methods(http.MethodPut)
	api.HandleFunc("/{account}/elastigroups/{elastigroup}", s.ElastigroupShowHandler).Methods(http.MethodGet)
	api.HandleFunc("/{account}/elastigroups/{elastigroup}", s.ElastigroupDeleteHandler).Methods(http.MethodDelete)

	// managedinstance handlers
	api.HandleFunc("/{account}/instances", s.ManagedInstanceListHandler).Methods(http.MethodGet)
	api.HandleFunc("/{account}/instances", s.ManagedInstanceCreateHandler).Methods(http.MethodPost)
	api.HandleFunc("/{account}/instances/{instance}", s.ManagedInstanceUpdateHandler).Methods(http.MethodPut)
	api.HandleFunc("/{account}/instances/{instance}", s.ManagedInstanceShowHandler).Methods(http.MethodGet)
	api.HandleFunc("/{account}/instances/{instance}", s.ManagedInstanceDeleteHandler).Methods(http.MethodDelete)
}
