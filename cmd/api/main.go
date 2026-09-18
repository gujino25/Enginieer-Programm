package main

import (
	"enginer/internal/repository"
	httptransport "enginer/internal/transport/http_transport"
	"fmt"
)

func main() {

	projectList := repository.NewProjectStore()
	httpHandlers := httptransport.NewProjectHandlers(projectList)
	httpServer := httptransport.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server", err)
	}
}
