package services

import "league/internal/repository"

type Services struct {
	Repos *repository.Repository
}

func New(repos *repository.Repository) *Services {
	return &Services{
		Repos: repos,
	}
}
