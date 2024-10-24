package service

import log "github.com/sirupsen/logrus"

type repo interface {
}

type Service struct {
	repo repo
}

func NewService(repo repo) *Service {
	log.Info("service init")

	return &Service{
		repo: repo,
	}
}
