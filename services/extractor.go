package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"
)

type (
	Extractor struct{}
)

var (
	initialized bool
	extractor   *Extractor
	logger      *zap.Logger

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
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = l
	extractor = NewExtractor()

	initialized = true
	logger.Info("Successfully initialized extractor")
}

func NewExtractor() *Extractor {
	return &Extractor{}
}

func (e *Extractor) Extract() error {
	return errNotImplemented
}
