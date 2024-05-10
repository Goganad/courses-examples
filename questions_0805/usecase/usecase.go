package usecase

import "miscellaneous/questions_0805/domain"

type UserRepository interface {
	GetByID(id string) (*domain.User, error)
	Save(user *domain.User) error
}

type UserManager struct {
	Repo UserRepository
}

func (m *UserManager) GetUser(id string) (*domain.User, error) {
	return m.Repo.GetByID(id)
}

func (m *UserManager) CreateUser(user *domain.User) error {
	return nil
}
