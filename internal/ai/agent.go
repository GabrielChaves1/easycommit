package ai

import "context"

type Agent interface {
	GenerateCommitMessage(ctx context.Context, diff, language string) (string, error)
}