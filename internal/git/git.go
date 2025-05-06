package git

import (
	"bytes"
	"errors"
	"os/exec"
)

var (
	ErrNoStagedChanges error = errors.New("no staged changes")
	ErrGitNotFound     error = errors.New("git not found or not a git repository")
)

func GetStagedDiff() (string, error) {
	// Execute the git diff command to get the staged changes
	cmd := exec.Command("git", "diff", "--cached", "--diff-algorithm=minimal")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return "", ErrNoStagedChanges
		}
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 128 {
			return "", ErrGitNotFound
		}
		return "", err
	}

	diff := out.String()
	if diff == "" {
		return "", ErrNoStagedChanges
	}

	return diff, nil
}

func CommitStagedChanges(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return "", err
	}

	return out.String(), nil
}
