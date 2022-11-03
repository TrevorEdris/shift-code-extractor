package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/bwmarrin/discordgo"
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type (
	Extractor struct {
		cfg            *config
		discordSession *discordgo.Session
		logger         *zap.Logger
	}

	config struct {
		Channel string `env:"CHANNEL_ID,required"`
		Token   string `env:"BOT_TOKEN,required"`
	}

	shiftCode struct {
		key     string
		game    string
		expires time.Time
	}
)

const (
	imageURL = `https://www.google.com/url?sa=i&url=https%3A%2F%2Fimgflip.com%2Fmemegenerator%2F255512681%2FBabe-Its-4PM-time-for-your&psig=AOvVaw0thNF3wktdXGoLEFlP8xbc&ust=1667528646255000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCJiW47b6kPsCFQAAAAAdAAAAABAE`
)

var (
	initialized bool

	extractor *Extractor

	errNotImplemented = errors.New("function not implemented")
)

func ExtractorHandler(ctx context.Context, event events.CloudWatchEvent) error {
	if !initialized {
		err := initialize()
		if err != nil {
			return err
		}
	}
	err := extractor.Extract()
	if err != nil {
		extractor.logger.Error("failed to extract keys", zap.Error(err))
		return fmt.Errorf("failed to extract keys: %w", err)
	}

	return nil
}

func initialize() error {
	l, err := initLogging()
	if err != nil {
		return err
	}
	cfg, err := initConfig()
	if err != nil {
		return err
	}
	extractor, err = initExtractor(cfg, l)
	if err != nil {
		return err
	}

	initialized = true
	l.Info("Successfully initialized extractor")

	return nil
}

func initLogging() (*zap.Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return l, nil
}

func initConfig() (*config, error) {
	// Attempt to load a .env file, but if it errors out and it's NOT a IsNotExist error, then
	// there was a problem parsing the .env file
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	// Initialize the cfg variable
	var cfg *config
	err = envdecode.StrictDecode(cfg)
	if err != nil {
		return nil, err
	}

	// Perform any custom validation steps
	err = cfg.Validate()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// probably doesn't _need_ to be a separate function, but it matches the pattern of the other funcs
func initExtractor(cfg *config, logger *zap.Logger) (*Extractor, error) {
	// Create a discord session using the bot token in the config variable
	dg, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		return nil, err
	}

	e := NewExtractor(cfg, dg, logger)
	return e, nil

}

func (c *config) Validate() error {
	return nil
}

func NewExtractor(cfg *config, session *discordgo.Session, logger *zap.Logger) *Extractor {
	return &Extractor{
		cfg:            cfg,
		discordSession: session,
		logger:         logger,
	}
}

func (e *Extractor) Extract() error {
	// Search last <duration> of tweets from [user1, user2, user3] for %pattern%
	e.logger.Info("Searching last <DURATION> of tweets from [USERS] for PATTERN")

	// Extract components of message
	e.logger.Info("Found matching tweet TWEET_ID with contents TWEET_CONTENTS")
	// TODO: Future PR, generate []shiftCode and send a message for each one
	codes := []shiftCode{
		{
			game:    "TODO - @MuchUsername implement this shit",
			key:     "abcdefgh-1234-567-9810-abcdefghijkl",
			expires: time.Now().Add(24 * time.Hour),
		},
	}

	// Send Discord message
	e.logger.Info("Sending messages to channel CHANNEL_ID", zap.Int("count", len(codes)))
	now := time.Now().Format(time.RFC3339)
	for _, code := range codes {
		msg, err := e.discordSession.ChannelMessageSendEmbed(
			e.cfg.Channel,
			&discordgo.MessageEmbed{
				Author:      &discordgo.MessageEmbedAuthor{},
				Color:       0x00ff00,
				Description: "Shift code",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:   "Game",
						Value:  code.game,
						Inline: true,
					},
					{
						Name:   "Key",
						Value:  code.key,
						Inline: true,
					},
					{
						Name:   "Expires",
						Value:  code.expires.Format(time.RFC3339),
						Inline: true,
					},
				},
				Image:     &discordgo.MessageEmbedImage{URL: imageURL},
				Thumbnail: &discordgo.MessageEmbedThumbnail{URL: imageURL},
				Timestamp: now,
				Title:     fmt.Sprintf("Babe! It's %s, time for your new Shift Code!", now),
			},
		)
		if err != nil {
			return err
		}
		e.logger.Info("Sent", zap.String("msg", msg.Content), zap.String("channelID", msg.ChannelID))
	}
	return nil
}
