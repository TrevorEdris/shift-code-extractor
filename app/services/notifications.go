package services

import (
	"github.com/TrevorEdris/shift-code-extractor/app/config"
	"github.com/TrevorEdris/shift-code-extractor/app/internal/repository"
)

type (
	NotificationSvc struct {
		repo      *repository.NotificationRepository
		publisher *repository.NotificationPublisher
	}
)

func NewNotificationSvc(cfg *config.Config) (*NotificationSvc, error) {
	return &NotificationSvc{}, nil
}
