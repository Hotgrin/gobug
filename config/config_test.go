package config

import "testing"

func TestSaveAndLoad(t *testing.T) {
	configDirOverride = t.TempDir()
	defer func() { configDirOverride = "" }()

	want := Config{APIKey: "sk-ant-test-key", Model: "claude-sonnet-4-5"}
	if err := Save(want); err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	got := Load()
	if got != want {
		t.Fatalf("Load() = %+v, want %+v", got, want)
	}
}

func TestLoad_NoConfigYet(t *testing.T) {
	configDirOverride = t.TempDir()
	defer func() { configDirOverride = "" }()

	got := Load()
	if got != (Config{}) {
		t.Fatalf("Load() with no saved config = %+v, want zero value", got)
	}
}
