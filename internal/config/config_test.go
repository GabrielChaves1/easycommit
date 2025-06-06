package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadNoFile(t *testing.T) {
	dir := t.TempDir()
	ConfigPath = filepath.Join(dir, "cfg.yaml")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(cfg, &Config{}) {
		t.Fatalf("expected empty config, got %+v", cfg)
	}
}

func TestSaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	ConfigPath = filepath.Join(dir, "cfg.yaml")

	original := &Config{AgentType: "openai", APIKey: "key", Language: "en"}
	if err := original.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("failed to load: %v", err)
	}

	if !reflect.DeepEqual(original, loaded) {
		t.Fatalf("mismatch after load: got %+v want %+v", loaded, original)
	}

	if _, err := os.Stat(ConfigPath); err != nil {
		t.Fatalf("config file not created: %v", err)
	}
}
