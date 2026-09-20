package main

import (
	"enginer/internal/repository"
	"enginer/internal/service"
	httptransport "enginer/internal/transport/http_transport"
	"fmt"
)

func main() {

	projectList := repository.NewProjectStore()
	systemList := repository.NewSystemStore()
	segmentList := repository.NewSegmentStore()
	segmentSvc := service.NewSegmentService(segmentList, systemList)
	systemSvc := service.NewSystemService(systemList, projectList)
	projectHandlers := httptransport.NewProjectHandlers(projectList)
	systemHadlers := httptransport.NewSystemHandlers(systemSvc)
	segmentHandlers := httptransport.NewSegmentHandlers(segmentSvc)
	httpServer := httptransport.NewHTTPServer(projectHandlers, systemHadlers, segmentHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server", err)
	}
}
