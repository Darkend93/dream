package repository

import (
	"context"
	"dream/domain/entities"
)

type DreamRepository interface {
	Insert(c context.Context, dm *entities.Dream) error
	FindAll(c context.Context) ([]*entities.Dream, error)
	FindById(c context.Context, id *int) (*entities.Dream, error)
	Update(c context.Context, de *entities.Dream) error
}
