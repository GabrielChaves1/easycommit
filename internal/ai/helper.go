package ai

import (
	"encoding/json"
	"fmt"
)

func GeneratePrompt(language string) string {
	return "You are an AI assistant specialized in creating concise and descriptive Git commit messages. " +
		"Given a git diff, create a clear commit message that follows best practices: " +
		"- Create ONLY a single-line commit title (no description body) " +
		"- Use the imperative mood (e.g., 'Add feature' not 'Added feature') " +
		"- Keep the message under 75 characters " +
		"- Be descriptive but concise " +
		"- Focus on WHY and WHAT, not HOW " +
		"- Assume the diff is a single commit" +
		"- Use the language " + language + " for the commit message." +
		BuildConventionalCommitPrompt()
}

func BuildConventionalCommitPrompt() string {
	types := map[string]string{
		"docs":     "Documentation only changes",
		"style":    "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)",
		"refactor": "A code change that neither fixes a bug nor adds a feature",
		"perf":     "A code change that improves performance",
		"test":     "Adding missing tests or correcting existing tests",
		"build":    "Changes that affect the build system or external dependencies",
		"ci":       "Changes to our CI configuration files and scripts",
		"chore":    "Other changes that don't modify src or test files",
		"revert":   "Reverts a previous commit",
		"feat":     "A new feature",
		"fix":      "A bug fix",
	}

	typesJSON, _ := json.MarshalIndent(types, "", "  ")

	return fmt.Sprintf(
		"Conventional Commit: Choose a type from the type-to-description JSON that best describes the git diff:\n%s\n",
		string(typesJSON),
	)
}
