package ai

import "github.com/GabrielChaves1/easycommit/internal/config"

type Option func(*config.Config)

func WithAPIKey(apiKey string) Option {
	return func(opts *config.Config) {
		opts.APIKey = apiKey
	}
}
