package httptransport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	projectHandlers *ProjectHandlers
	systemHandlers  *SystemHandlers
}

func NewHTTPServer(projectHandler *ProjectHandlers, systemHadlers *SystemHandlers) *HTTPServer {
	return &HTTPServer{
		projectHandlers: projectHandler,
		systemHandlers:  systemHadlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()
	router.Path("/projects").Methods("POST").HandlerFunc(s.projectHandlers.HandleCreateProject)
	router.Path("/projects/{id}").Methods("GET").HandlerFunc(s.projectHandlers.HandleGetProject)
	router.Path("/projects").Methods("GET").HandlerFunc(s.projectHandlers.HandleListProjects)
	router.Path("/projects/{project_id}/systems").Methods("POST").HandlerFunc(s.systemHandlers.HandleCreateSystem)
	router.Path("/projects/{project_id}/systems").Methods("GET").HandlerFunc(s.systemHandlers.HandleListByProject)
	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
