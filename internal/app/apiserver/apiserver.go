package apiserver

import "github.com/gorilla/mux"

type APIServer struct {
	router *mux.Router
}
