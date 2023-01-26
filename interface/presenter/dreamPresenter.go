package presenter

import (
	"dream/models"
	"dream/usecases/presenter"
)

type dreamPresenter struct{}

func NewDreamPresenter() presenter.DreamPresenter {
	return &dreamPresenter{}
}

func (p *dreamPresenter) ResponseDreams(dm []*models.DreamModel) []*models.DreamModel {
	return dm
}
