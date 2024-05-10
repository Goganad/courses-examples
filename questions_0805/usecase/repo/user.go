package repo

import "miscellaneous/questions_0805/domain"

type InMemoryUserRepository struct {
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (repo *InMemoryUserRepository) GetByID(id string) (*domain.User, error) {
	if user, ok := repo.users[id]; ok {
		return user, nil
	}
	return nil, nil
}

func (repo *InMemoryUserRepository) Save(user *domain.User) error {
	repo.users[user.ID] = user
	return nil
}
