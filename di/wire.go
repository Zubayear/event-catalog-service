//go:build wireinject
// +build wireinject

package di

import (
	"EventCatalog/external"
	"EventCatalog/handler"
	"EventCatalog/repository"
	"github.com/google/wire"
)

func DependencyProvider() (*handler.EventCatalog, error) {
	wire.Build(repository.DatabaseImlProvider, handler.NewEventCatalog, external.HostProvider,
		wire.Bind(new(repository.IRepository), new(*repository.DatabaseImpl)))
	return &handler.EventCatalog{}, nil
}
