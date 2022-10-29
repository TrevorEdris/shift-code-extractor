package services

import (
	"github.com/TrevorEdris/shift-code-extractor/app/config"
	"github.com/TrevorEdris/shift-code-extractor/app/domain"
	"github.com/TrevorEdris/shift-code-extractor/app/internal/repository"
)

type (
	TweetSvc struct {
		repo   repository.TweetRepository
		client repository.TwitterClient
	}
)

func NewTweetSvc(cfg *config.Config) (*TweetSvc, error) {
	return &TweetSvc{}, nil
}

func (t *TweetSvc) GetRecentTweets() ([]domain.Tweet, error) {
	return nil, errNotImplemented
}
