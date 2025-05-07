package git

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

var (
	ErrNoStagedChanges = errors.New("no staged changes")
	ErrGitNotFound     = errors.New("git not found or not a git repository")
)

func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached", "--diff-algorithm=minimal")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()

	// If the command fails, check if it's due to not being in a git repository
	if err != nil {
		if strings.Contains(out.String(), "Not a git repository") {
			return "", ErrGitNotFound
		}
		return "", err
	}

	// Check if there are any staged changes
	diff := out.String()
	if strings.TrimSpace(diff) == "" {
		return "", ErrNoStagedChanges
	}

	return diff, nil
}

func CommitStagedChanges(message string) (string, error) {
	// Sanitize the message to avoid issues with newlines
	safeMsg := strings.ReplaceAll(message, "\n", " ")
	cmd := exec.Command("git", "commit", "-m", safeMsg)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		// If the error is due to no staged changes
		if strings.Contains(stderr.String(), "nothing to commit") {
			return "", ErrNoStagedChanges
		}
		return "", errors.New(stderr.String())
	}

	return out.String(), nil
}