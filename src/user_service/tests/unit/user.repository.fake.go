package unit

import (
	"User_Service/domain"
)

type InMemoryUserRepository struct {
	users []domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	repo := &InMemoryUserRepository{}

	repo.users = []domain.User{{Email: "test_email@abc.com", Password: "1234"},
		{Email: "ben.parker@ny.com", Password: "benP"}}

	return repo
}

func (repo *InMemoryUserRepository) Add(user domain.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *InMemoryUserRepository) Get(email string) (*domain.User, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
}
