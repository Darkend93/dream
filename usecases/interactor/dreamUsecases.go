package interactor

import (
	"context"
	"dream/domain/entities"
	m "dream/domain/models"
	"dream/usecases/presenter"
	"dream/usecases/repository"
)

type dreamInteractor struct {
	DreamRepository repository.DreamRepository
	DreamPresenter  presenter.DreamPresenter
}

type DreamInteractor interface {
	Add(c context.Context, dm *m.DreamModel) error
	GetAll(c context.Context) ([]*m.DreamModel, error)
	GetById(c context.Context, id *int) (*m.DreamModel, error)
}

func NewDreamInteractor(r repository.DreamRepository, p presenter.DreamPresenter) DreamInteractor {
	return &dreamInteractor{r, p}
}

func (i *dreamInteractor) Add(c context.Context, dm *m.DreamModel) error {
	if err := i.DreamRepository.Insert(c, &entities.Dream{
		Name:        dm.Name,
		Description: dm.Description,
	}); err != nil {
		return err
	}

	return nil
}

func (i *dreamInteractor) GetAll(c context.Context) ([]*m.DreamModel, error) {
	var result []*m.DreamModel

	dreams, err := i.DreamRepository.FindAll(c)
	if err != nil {
		return nil, err
	}

	for _, d := range dreams {
		result = append(result, &m.DreamModel{
			Name:        d.Name,
			Description: d.Description,
		})
	}

	return result, nil
}

func (i *dreamInteractor) GetById(c context.Context, id *int) (*m.DreamModel, error) {
	d, err := i.DreamRepository.FindById(c, id)
	if err != nil {
		return nil, err
	}

	return &m.DreamModel{
		Name:        d.Name,
		Description: d.Description,
	}, nil
}
