package presenter

import "dream/domain/models"

type DreamPresenter interface {
	ResponseDreams(d []*models.DreamModel) []*models.DreamModel
}
