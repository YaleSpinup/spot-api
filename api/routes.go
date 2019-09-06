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
	api.HandleFunc("/{account}/elastigroups/{elastigroup}", s.ElastigroupShowHandler).Methods(http.MethodGet)
	api.HandleFunc("/{account}/elastigroups/{elastigroup}", s.ElastigroupDeleteHandler).Methods(http.MethodDelete)
	// api.HandleFunc("/{account}/elastigroups/{elastigroup}", s.ElastigrouptUpdateHandler).Methods(http.MethodPut)
}
