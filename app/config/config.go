package config

import (
	"errors"
	"os"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App
		AWS
		OAuth
		Tweets
		Notifications
	}

	App struct {
		Name string `env:"APP_NAME,default=shift-code-extractor"`
	}

	AWS struct {
		AccessKey string `env:"AWS_ACCESS_KEY_ID"`
		Secret    string `env:"AWS_SECRET_ACCESS_KEY"`
	}

	OAuth struct {
		Token  string `env:"TWITTER_OAUTH_TOKEN"`
		Secret string `env:"TWITTER_OAUTH_SECRET"`
	}

	Tweets struct {
		SourceUsernames []string `env:"TWEET_SOURCE_USERNAMES"`
		Storage         string   `env:"TWEET_STORAGE_TYPE,default=local"`
	}

	Notifications struct {
		Storage string `env:"NOTIFICATION_STORAGE_TYPE,default=local"`
	}

	DynamoDB struct {
		TweetsTable        string `env:"DYNAMODB_TWEETS_TABLE"`
		NotificationsTable string `env:"DYNAMODB_NOTIFICATIONS_TABLE"`
	}

	SNS struct {
		Topic string `env:"SNS_TOPIC"`
	}
)

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	cfg := Config{}
	err = envdecode.StrictDecode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
