package main

import (
	ec "EventCatalog/proto"
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go-micro.dev/v4/client"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var (
	gateway = flag.Int("gateway", 8081, "gateway")
	port    = flag.Int("port", 60007, "eventcatalog address")
	gRPC    = flag.Int("gRPC", 60007, "gRPC port")
)

var eventCatalogClient ec.EventCatalogService

func init() {
	eventCatalogClient = ec.NewEventCatalogService("eventcatalog", client.DefaultClient)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	s := grpc.NewServer()

	ec.RegisterEventCatalogServer(s, &EventCatalogProxy{})

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		return err
	}

	go func() {
		err := s.Serve(listen)
		if err != nil {
			return
		}
	}()

	mux := runtime.NewServeMux()
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", *gRPC), grpc.WithInsecure())
	if err != nil {
		return err
	}

	err = ec.RegisterEventCatalogEC(ctx, mux, conn)
	if err != nil {
		return err
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", *gateway),
		Handler: mux,
	}

	return gwServer.ListenAndServe()
}

func main() {
	flag.Parse()

	defer glog.Flush()
	err := run()
	if err != nil {
		return
	}
}

type EventCatalogProxy struct {
}

func (e *EventCatalogProxy) CreateEventWithCategory(ctx context.Context, request *ec.CreateEventWithCategoryRequest) (*ec.CreateEventWithCategoryResponse, error) {
	eventToReturn, err := eventCatalogClient.CreateEventWithCategory(ctx, request)
	if err != nil {
		return nil, err
	}
	return eventToReturn, nil
}

func (e *EventCatalogProxy) GetAllEvents(ctx context.Context, request *ec.GetAllEventsRequest) (*ec.GetAllEventsResponse, error) {
	eventToReturn, err := eventCatalogClient.GetAllEvents(ctx, request)
	if err != nil {
		return nil, err
	}
	return eventToReturn, nil
}

func (e *EventCatalogProxy) GetEventById(ctx context.Context, request *ec.GetEventByIdRequest) (*ec.GetEventByIdResponse, error) {
	eventToReturn, err := eventCatalogClient.GetEventById(ctx, request)
	if err != nil {
		return nil, err
	}
	return eventToReturn, nil
}

func (e *EventCatalogProxy) GetAllCategories(ctx context.Context, request *ec.GetAllCategoriesRequest) (*ec.GetAllCategoriesResponse, error) {
	eventToReturn, err := eventCatalogClient.GetAllCategories(ctx, request)
	if err != nil {
		return nil, err
	}
	return eventToReturn, nil
}

func (e *EventCatalogProxy) UpdateEvent(ctx context.Context, request *ec.UpdateEventByEventIdRequest) (*ec.UpdateEventByEventIdResponse, error) {
	eventToReturn, err := eventCatalogClient.UpdateEvent(ctx, request)
	if err != nil {
		return nil, err
	}
	return eventToReturn, nil
}

func (e *EventCatalogProxy) DeleteEventById(ctx context.Context, request *ec.DeleteEventByIdRequest) (*ec.DeleteEventByIdResponse, error) {
	eventToReturn, err := eventCatalogClient.DeleteEventById(ctx, request)
	if err != nil {
		return nil, err
	}
	return eventToReturn, nil
}

func (e *EventCatalogProxy) GetEventByField(ctx context.Context, request *ec.GetEventByFieldRequest) (*ec.GetEventByFieldResponse, error) {
	eventToReturn, err := eventCatalogClient.GetEventByField(ctx, request)
	if err != nil {
		return nil, err
	}
	return eventToReturn, nil
}

func (e *EventCatalogProxy) CreateCategory(ctx context.Context, request *ec.CreateCategoryRequest) (*ec.CreateCategoryResponse, error) {
	eventToReturn, err := eventCatalogClient.CreateCategory(ctx, request)
	if err != nil {
		return nil, err
	}
	return eventToReturn, nil
}
