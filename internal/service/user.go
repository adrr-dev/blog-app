package service

import "github.com/adrr-dev/blog-app/internal/domain"

func (s Service) NewUser(username, password string) error {
	err := s.repo.CreateUser(username, password)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) FetchUser(username, password string) (*domain.User, error) {
	user, err := s.repo.FetchUser(username, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s Service) FetchUserByID(id uint) (*domain.User, error) {
	user, err := s.repo.FetchUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
