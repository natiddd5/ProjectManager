package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr  string
	store Store
}

func newAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

// initalize the router and register all services and their dependecies.
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//register all services
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(subrouter)

	userService := NewUserService(s.store)
	userService.RegisterRoutes(subrouter)

	projectService := NewProjectService(s.store)
	projectService.RegisterRoutes(subrouter)

	log.Println("Starting the server at ", s.addr)
	log.Fatalln(http.ListenAndServe(s.addr, subrouter))
}
