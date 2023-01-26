package registry

import (
	"dream/interface/handlers"
	ipres "dream/interface/presenter"
	irepo "dream/interface/repository"
	"dream/usecases/interactor"
	pres "dream/usecases/presenter"
	repo "dream/usecases/repository"
)

func (r *registry) NewDreamHandler() handlers.DreamHandler {
	return handlers.NewDreamHandler(r.NewDreamInteractor())
}

func (r *registry) NewDreamInteractor() interactor.DreamInteractor {
	return interactor.NewDreamInteractor(r.NewDreamRepository(), r.NewDreamPresenter())
}

func (r *registry) NewDreamRepository() repo.DreamRepository {
	return irepo.NewDreamRepository(r.db)
}

func (r *registry) NewDreamPresenter() pres.DreamPresenter {
	return ipres.NewDreamPresenter()
}
