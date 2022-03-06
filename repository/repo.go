package repository

import (
	"EventCatalog/ent"
	"EventCatalog/ent/category"
	"EventCatalog/ent/migrate"
	"EventCatalog/external"
	pb "EventCatalog/proto"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	log "go-micro.dev/v4/logger"
	"time"
)

type IRepository interface {
	Update(ctx context.Context, entity interface{}) (interface{}, error)
	Delete(ctx context.Context, entity interface{}) (interface{}, error)
	GetByField(ctx context.Context, field interface{}) (interface{}, error)
	SaveParent(ctx context.Context, entity interface{}) (interface{}, error)
	GetAllParent(ctx context.Context) (interface{}, error)
	SaveChildWithParent(ctx context.Context, parent, child interface{}) (interface{}, error)
	GetChildById(ctx context.Context, entity interface{}) (interface{}, error)
	GetAllChild(ctx context.Context) (interface{}, error)
}

type DatabaseImpl struct {
	client *ent.Client
}

func DatabaseImlProvider(h *external.Host) (*DatabaseImpl, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", h.User, h.Password, h.Address, h.Port, h.Name)
	//client, err := ent.Open(h.Type, "root:root@tcp(localhost:3306)/EventCatalogDB?parseTime=true")
	client, err := ent.Open(h.Type, connString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	//defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return &DatabaseImpl{client: client}, nil
}

func (d *DatabaseImpl) GetAllChild(ctx context.Context) (interface{}, error) {
	allEvents, err := d.client.Event.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting all events: %w", err)
	}
	var eventsToReturn []*pb.Event
	for _, event := range allEvents {
		eventFromRepo := &pb.Event{}
		eventMapper(event, eventFromRepo, false)
		eventsToReturn = append(eventsToReturn, eventFromRepo)
	}
	return eventsToReturn, nil
}

func (d *DatabaseImpl) GetChildById(ctx context.Context, entity interface{}) (interface{}, error) {
	id, err := uuid.Parse(entity.(string))
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}
	eventFromRepo, err := d.client.Event.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed getting event: %w", err)
	}
	eventToReturn := &pb.Event{}
	eventMapper(eventFromRepo, eventToReturn, false)
	return eventToReturn, nil
}

func (d *DatabaseImpl) SaveChildWithParent(ctx context.Context, parent, child interface{}) (interface{}, error) {
	categoryId := parent.(string)
	eventToSave := child.(*pb.Event)
	categoryFromRepo, err := d.client.Category.Get(ctx, uuid.MustParse(categoryId))
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}
	savedEvent, err := d.client.Event.Create().
		SetName(eventToSave.Name).
		SetPrice(eventToSave.Price).
		SetDescription(eventToSave.Description).
		SetDate(time.Now().Unix()).
		SetArtist(eventToSave.Artist).
		SetImageUrl(eventToSave.ImageUrl).
		SetOwner(categoryFromRepo).Save(ctx)
	if err != nil {
		return nil, err
	}
	eventToReturn := &pb.Event{}
	eventMapper(savedEvent, eventToReturn, false)
	return eventToReturn, nil
}

func (d *DatabaseImpl) GetAllParent(ctx context.Context) (interface{}, error) {
	categories, err := d.client.Category.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var categoriesToReturn []string
	for _, c := range categories {
		categoriesToReturn = append(categoriesToReturn, c.Name.String())
	}
	return categoriesToReturn, nil
}

func (d *DatabaseImpl) SaveParent(ctx context.Context, entity interface{}) (interface{}, error) {
	categoryToSave := entity.(string)
	savedCategory, err := d.client.Category.Create().SetName(category.Name(categoryToSave)).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating category: %w", err)
	}
	return savedCategory.Name.String(), nil
}

func (d *DatabaseImpl) GetByField(ctx context.Context, field interface{}) (interface{}, error) {
	eventsFromRepo, err := d.client.Event.
		Query().Select(field.(string)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed getting event: %w", err)
	}
	var eventsToReturn []*pb.Event
	for _, event := range eventsFromRepo {
		eventFromRepo := &pb.Event{}
		eventMapper(event, eventFromRepo, false)
		eventsToReturn = append(eventsToReturn, eventFromRepo)
	}
	return eventsToReturn, nil
}

func (d *DatabaseImpl) Delete(ctx context.Context, entity interface{}) (interface{}, error) {
	id := entity.(string)
	err := d.client.Event.DeleteOneID(uuid.MustParse(id)).Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed deleting event: %w", err)
	}
	return "event deleted", nil
}

func (d *DatabaseImpl) Update(ctx context.Context, entity interface{}) (interface{}, error) {
	eventToUpdate := entity.(*pb.UpdateEventByEventIdRequest)
	id := uuid.MustParse(eventToUpdate.EventId)
	updatedEvent, err := d.client.Event.
		UpdateOneID(id).
		SetName(eventToUpdate.GetName()).
		SetPrice(eventToUpdate.GetPrice()).
		SetArtist(eventToUpdate.GetArtist()).
		SetDescription(eventToUpdate.GetDescription()).
		SetImageUrl(eventToUpdate.GetImageUrl()).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	eventToReturn := &pb.Event{}
	eventMapper(updatedEvent, eventToReturn, false)
	return eventToReturn, nil
}

func (d *DatabaseImpl) Save(ctx context.Context, entity interface{}) (interface{}, error) {
	event := entity.(*pb.Event)
	savedEvent, err := d.client.Event.Create().
		SetName(event.Name).
		SetPrice(event.Price).
		SetArtist(event.Artist).
		SetDate(time.Now().Unix()).
		SetDescription(event.Description).
		SetImageUrl(event.ImageUrl).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating event: %w", err)
	}
	eventToReturn := &pb.Event{}
	eventMapper(savedEvent, eventToReturn, false)
	return eventToReturn, nil
}

func eventMapper(src *ent.Event, dst *pb.Event, rev bool) {
	switch rev {
	case true:
		src.ID = uuid.MustParse(dst.Id)
		src.Name = dst.Name
		src.Description = dst.Description
		src.Price = dst.Price
		src.Artist = dst.Artist
		src.Date = dst.Date
		src.ImageUrl = dst.ImageUrl
	case false:
		dst.Id = src.ID.String()
		dst.Name = src.Name
		dst.Description = src.Description
		dst.Price = src.Price
		dst.Artist = src.Artist
		dst.Date = src.Date
		dst.ImageUrl = src.ImageUrl
	}

}

//func categoryMapper(src *ent.Category, dst *pb.Category) {
//	src.Name = dst.String()
//}

func NewDatabaseImpl() IRepository {
	return &DatabaseImpl{}
}
