package handler

import (
	pb "EventCatalog/proto"
	"EventCatalog/repository"
	"context"
	log "go-micro.dev/v4/logger"
)

type EventCatalog struct {
	repo repository.IRepository
}

func NewEventCatalog(r repository.IRepository) (*EventCatalog, error) {
	return &EventCatalog{repo: r}, nil
}

func (e *EventCatalog) CreateEventWithCategory(ctx context.Context, request *pb.CreateEventWithCategoryRequest, response *pb.CreateEventWithCategoryResponse) error {
	log.Infof("Received EventCatalog.CreateEventWithCategory request: %v", request)
	categoryToSave := request.CategoryId
	eventToSave := &pb.Event{
		Name:        request.Name,
		Price:       request.Price,
		Artist:      request.Artist,
		Date:        request.Date,
		Description: request.Description,
		ImageUrl:    request.ImageUrl,
	}
	eventToReturn, err := e.repo.SaveChildWithParent(ctx, categoryToSave, eventToSave)
	if err != nil {
		return err
	}
	response.Event = eventToReturn.(*pb.Event)
	return nil
}

func (e *EventCatalog) CreateCategory(ctx context.Context, request *pb.CreateCategoryRequest, response *pb.CreateCategoryResponse) error {
	log.Infof("Received EventCatalog.CreateCategory request: %v", request)
	// map here so db don't have to give validator failed for field "Category.Name"
	categoryFromRepo, err := e.repo.SaveParent(ctx, request.GetCategory().String())
	if err != nil {
		return err
	}
	response.Category = categoryFromRepo.(string)
	return nil
}

func (e *EventCatalog) GetEventByField(ctx context.Context, request *pb.GetEventByFieldRequest, response *pb.GetEventByFieldResponse) error {
	log.Infof("Received EventCatalog.GetEventByField request: %v", request)
	eventToReturn, err := e.repo.GetByField(ctx, request.GetFieldName())
	if err != nil {
		return err
	}
	response.Events = eventToReturn.([]*pb.Event)
	return nil
}

func (e *EventCatalog) GetEventBySpec(ctx context.Context, request *pb.GetEventBySpecRequest, response *pb.GetEventBySpecResponse) error {
	log.Infof("Received EventCatalog.DeleteEventById request: %v", request)
	eventToReturn, err := e.repo.GetByField(ctx, request.GetFieldName())
	if err != nil {
		return err
	}
	response.Events = eventToReturn.([]*pb.Event)
	return nil
}

func (e *EventCatalog) DeleteEventById(ctx context.Context, request *pb.DeleteEventByIdRequest, response *pb.DeleteEventByIdResponse) error {
	log.Infof("Received EventCatalog.DeleteEventById request: %v", request)
	msg, err := e.repo.Delete(ctx, request.GetId())
	if err != nil {
		return err
	}
	response.Message = msg.(string)
	return nil
}

func (e *EventCatalog) UpdateEvent(ctx context.Context, request *pb.UpdateEventByEventIdRequest, response *pb.UpdateEventByEventIdResponse) error {
	log.Infof("Received EventCatalog.UpdateEvent request: %v", request)
	updatedEventToReturn, err := e.repo.Update(ctx, request)
	if err != nil {
		return err
	}
	response.Event = updatedEventToReturn.(*pb.Event)
	return nil
}

func (e *EventCatalog) GetAllEvents(ctx context.Context, request *pb.GetAllEventsRequest, response *pb.GetAllEventsResponse) error {
	log.Infof("Received EventCatalog.GetAllEvents request: %v", request)
	eventsFromRepo, err := e.repo.GetAllChild(ctx)
	if err != nil {
		return err
	}
	response.Events = eventsFromRepo.([]*pb.Event)
	return nil
}

func (e *EventCatalog) GetEventById(ctx context.Context, request *pb.GetEventByIdRequest, response *pb.GetEventByIdResponse) error {
	log.Infof("Received EventCatalog.GetEventById request: %v", request)
	event, err := e.repo.GetChildById(ctx, request.GetId())
	if err != nil {
		return err
	}
	response.Event = event.(*pb.Event)
	return nil
}

func (e *EventCatalog) GetAllCategories(ctx context.Context, request *pb.GetAllCategoriesRequest, response *pb.GetAllCategoriesResponse) error {
	log.Infof("Received EventCatalog.GetAllCategories request: %v", request)
	categoriesToReturn, err := e.repo.GetAllParent(ctx)
	if err != nil {
		return err
	}
	response.Categories = categoriesToReturn.([]string)
	return nil
}

func NewEventCatalogHandler() pb.EventCatalogHandler {
	return &EventCatalog{}
}
