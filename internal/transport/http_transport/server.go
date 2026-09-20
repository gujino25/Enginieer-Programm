package httptransport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	projectHandlers *ProjectHandlers
	systemHandlers  *SystemHandlers
	segmentHandlers *SegmentHandlers
}

func NewHTTPServer(
	projectHandler *ProjectHandlers,
	systemHadlers *SystemHandlers,
	segmentHandlers *SegmentHandlers) *HTTPServer {
	return &HTTPServer{
		projectHandlers: projectHandler,
		systemHandlers:  systemHadlers,
		segmentHandlers: segmentHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()
	router.Path("/projects").Methods("POST").HandlerFunc(s.projectHandlers.HandleCreateProject)
	router.Path("/projects/{id}").Methods("GET").HandlerFunc(s.projectHandlers.HandleGetProject)
	router.Path("/projects").Methods("GET").HandlerFunc(s.projectHandlers.HandleListProjects)
	router.Path("/projects/{project_id}/systems").Methods("POST").HandlerFunc(s.systemHandlers.HandleCreateSystem)
	router.Path("/projects/{project_id}/systems").Methods("GET").HandlerFunc(s.systemHandlers.HandleListByProject)
	router.Path("/systems/{system_id}/segments").Methods("POST").HandlerFunc(s.segmentHandlers.HandleCreateSegment)
	router.Path("/systems/{system_id}/segments").Methods("GET").HandlerFunc(s.segmentHandlers.HandleListBySystem)
	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
