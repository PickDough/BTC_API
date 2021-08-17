package unit

import "SE_School/models"

type InMemoryUserRepository struct {
	users []models.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	repo := &InMemoryUserRepository{}

	repo.users = []models.User{{Email: "test_email@abc.com", Password: "1234"},
		{Email: "ben.parker@ny.com", Password: "benP"}}

	return repo
}

func (repo *InMemoryUserRepository) Add(user models.User) error {
	repo.users = append(repo.users, user)
	return nil
}

func (repo *InMemoryUserRepository) Get(email string) (*models.User, error) {
	for _, user := range repo.users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, nil
}
