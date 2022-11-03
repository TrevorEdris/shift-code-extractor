package services

import (
	"context"
	"errors"
	"fmt"
	"os"

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
		key string
		game string
		expires time.Time
	}
)

var (
	initialized bool

	extractor *Extractor

	errNotImplemented = errors.New("function not implemented")
)

func ExtractorHandler(ctx context.Context, event events.CloudWatchEvent) error {
	if !initialized {
		initialize()
	}
	err := extractor.Extract()
	if err != nil {
		logger.Error("Error extracting keys", zap.Error(err))
		return fmt.Errorf("failed to extract keys: %w", err)
	}

	return nil
}

func initialize() {
	l := initLogging()
	cfg := initConfig()
	initExtractor(cfg)

	initialized = true
	logger.Info("Successfully initialized extractor")
}

func initLogging() *zap.Logger {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return l
}

func initConfig() *config {
	// Attempt to load a .env file, but if it errors out and it's NOT a IsNotExist error, then
	// there was a problem parsing the .env file
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	// Initialize the cfg variable
	var cfg *config
	err = envdecode.StrictDecode(cfg)
	if err != nil {
		panic(err)
	}

	// Perform any custom validation steps
	err = cfg.Validate()
	if err != nil {
		panic(err)
	}

	return cfg
}

// probably doesn't _need_ to be a separate function, but it matches the pattern of the other funcs
func initExtractor(cfg *config, logger *zap.Logger) {
	// Create a discord session using the bot token in the config variable
	dg, err := discordgo.new("Bot " + cfg.Token)
	if err != nil {
		panic(err)
	}

	extractor = NewExtractor(cfg, logger)
}

func (c *config) Validate() error {
	return nil
}

func NewExtractor(cfg *config, logger *zap.Logger) *Extractor {
	return &Extractor{
		cfg:    cfg,
		logger: logger,
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
			game: "TODO - @MuchUsername implement this shit",
			key: "abcdefgh-1234-567-9810-abcdefghijkl",
			expires: time.Now().Add(time.Day),
		},
	}

	// Send Discord message
	e.logger.Info("Sending messages to channel CHANNEL_ID", zap.Int("count", len(codes)))
	now := time.Now().Format(time.RFC3339)
	for _, code := range codes {
		e.discordSession.ChannelMessageSendEmbed(
			e.cfg.Channel,
			&discordgo.MessageEmbed{
				Author: &discordgo.MessageEmbedAuthor{},
				Color: 0x00ff00, // Green
				Description: "Shift code",
				Fields []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name: "Game",
						Value: code.game,
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name: "Key",
						Value: code.key,
						Inline: true,
					},
					&discordgo.MessageEmbedField{
						Name: "Expires",
						Value: code.expires.Format(time.RFC3339),
						Inline: true,
					},
				},
			},
			Image: &discordgo.MessageEmbedImage{
				// TODO: I'd be surprised if this works lmao
				URL: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fimgflip.com%2Fmemegenerator%2F255512681%2FBabe-Its-4PM-time-for-your&psig=AOvVaw0thNF3wktdXGoLEFlP8xbc&ust=1667528646255000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCJiW47b6kPsCFQAAAAAdAAAAABAE",
			},
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: "https://www.google.com/url?sa=i&url=https%3A%2F%2Fimgflip.com%2Fmemegenerator%2F255512681%2FBabe-Its-4PM-time-for-your&psig=AOvVaw0thNF3wktdXGoLEFlP8xbc&ust=1667528646255000&source=images&cd=vfe&ved=0CAwQjRxqFwoTCJiW47b6kPsCFQAAAAAdAAAAABAE",
			},
			Timestamp: now,
			Title: fmt.Sprintf("Babe! It's %s, time for your new Shift Code!", now),
		)
	}
	return errNotImplemented
}
