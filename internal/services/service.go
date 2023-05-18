package services

import "league/internal/repositories"

type Services struct {
	Repos *repositories.Repository
}

func New(repos *repositories.Repository) *Services {
	return &Services{
		Repos: repos,
	}
}
