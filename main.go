package main

import (
	"EventCatalog/di"
	pb "EventCatalog/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "eventcatalog"
	version = "latest"
	port    = ":60007"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
		micro.Address(port),
	)
	srv.Init()

	h, _ := di.DependencyProvider()

	// Register handler
	err := pb.RegisterEventCatalogHandler(srv.Server(), h)
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
