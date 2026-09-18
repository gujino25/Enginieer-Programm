package httptransport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *ProjectHandlers
}

func NewHTTPServer(httpHandler *ProjectHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()
	router.Path("/projects").Methods("POST").HandlerFunc(s.httpHandlers.HandleCreateProject)
	router.Path("/projects/{id}").Methods("GET").HandlerFunc(s.httpHandlers.HandleGetProject)
	router.Path("/projects").Methods("GET").HandlerFunc(s.httpHandlers.HandleListProjects)
	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
