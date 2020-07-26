package usecase

import "go-clean-architecture/src/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) error {
	_, err := interactor.UserRepository.Store(u)
	return err
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
