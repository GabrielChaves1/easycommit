package ai

import (
	"strings"
	"testing"
)

func TestGeneratePromptContainsLanguage(t *testing.T) {
	prompt := GeneratePrompt("pt-BR")
	if !strings.Contains(prompt, "pt-BR") {
		t.Errorf("language not found in prompt")
	}
	if !strings.Contains(prompt, "Conventional Commit") {
		t.Errorf("conventional commit guidance missing")
	}
}

func TestBuildConventionalCommitPrompt(t *testing.T) {
	p := BuildConventionalCommitPrompt()
	if !strings.Contains(p, "Conventional Commit") {
		t.Errorf("prompt missing explanation")
	}
	if !strings.Contains(p, "feat") {
		t.Errorf("types json not included")
	}
}
