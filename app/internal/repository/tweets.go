package repository

import "github.com/TrevorEdris/shift-code-extractor/app/domain"

type (
	TweetRepository interface {
		Store(tweets []domain.Tweet) error
	}

	TwitterClient interface {
		GetRecentTweets(username string, limit int) ([]domain.Tweet, error)
	}

	NotificationRepository interface {
		Store(notifications []domain.Notification) error
	}

	NotificationPublisher interface {
		Publish(notification domain.Notification) error
	}
)
