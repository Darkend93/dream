package registry

import (
	"dream/interface/handlers"

	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppHandler() handlers.AppHandler
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppHandler() handlers.AppHandler {
	return handlers.AppHandler{
		Dream: r.NewDreamHandler(),
	}
}
