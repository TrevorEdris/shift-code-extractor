package services

import (
	"errors"
	"fmt"

	"github.com/TrevorEdris/shift-code-extractor/app/config"
)

var (
	errNotImplemented = errors.New("not implemented")
)

type (
	App struct {
		cfg *config.Config
	}
)

func NewApp() (*App, error) {
	a := App{}
	err := a.init()
	if err != nil {
		return nil, fmt.Errorf("failed to init app: %w", err)
	}

	return &a, nil
}

func (a *App) init() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	a.cfg = cfg

	return nil
}

func (a *App) Run() error {
	fmt.Println("Successfully ran")
	return nil
}
