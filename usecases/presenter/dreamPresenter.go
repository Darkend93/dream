package presenter

import "dream/models"

type DreamPresenter interface {
	ResponseDreams(d []*models.DreamModel) []*models.DreamModel
}
