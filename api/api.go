package api

import (
	"database/sql"
	"net/http"
)

type ApiServer struct {
	listenAddr string
	db         *sql.DB
	mux        *http.ServeMux
}

func NewApiServer(listenAddr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		db:         db,
		mux:        http.NewServeMux(),
	}
}

func (s *ApiServer) Start() error {
	return http.ListenAndServe(s.listenAddr, s.mux)
}

func (s *ApiServer) Get(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, ApiGet(s, handler))
}

func (s *ApiServer) Post(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, ApiPost(s, handler))
}

func (s *ApiServer) RegisterService(service ApiService) {
	service.RegisterRoutes(s)
}
