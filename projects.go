package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type ProjectService struct {
	store Store
}

func NewProjectService(s Store) *ProjectService {
	return &ProjectService{store: s}
}

func (s *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", s.handleCreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", s.handleGetProject).Methods("GET")
	r.HandleFunc("/projects/{id}", s.handleDeleteProject).Methods("DELETE")

}

// TODO:fix bug where project ID is written back as 0
func (s *ProjectService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var project *Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if project.Name == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Name is required"})
		return
	}

	err = s.store.CreateProject(project)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating project"})
		return
	}

	WriteJSON(w, http.StatusCreated, project)
}

func (s *ProjectService) handleGetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	project, err := s.store.GetProject(id)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error getting project"})
		return
	}

	WriteJSON(w, http.StatusOK, project)
}

func (s *ProjectService) handleDeleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := s.store.DeleteProject(id)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error deleting project"})
		return
	}

	WriteJSON(w, http.StatusNoContent, nil)
}
