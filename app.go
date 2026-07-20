package main

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"gobug/config"
	"gobug/engine"
)

// App holds the Wails runtime context and is bound to the frontend.
type App struct {
	ctx context.Context
	cfg config.Config
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.cfg = config.Load()
}

// GetConfig returns the current BYOK settings so the frontend can prefill
// the settings modal.
func (a *App) GetConfig() config.Config {
	return a.cfg
}

// SaveConfig persists BYOK settings locally. Returns an error message
// string (empty on success) since Wails bindings handle plain returns
// more simply than Go's (value, error) pattern on the JS side.
func (a *App) SaveConfig(apiKey string, model string) string {
	a.cfg = config.Config{APIKey: apiKey, Model: model}
	if err := config.Save(a.cfg); err != nil {
		return err.Error()
	}
	return ""
}

// RunResult is sent back to the frontend after a run attempt.
type RunResult struct {
	Stdout        string `json:"stdout"`
	Stderr        string `json:"stderr"`
	Success       bool   `json:"success"`
	Explanation   string `json:"explanation"`
	ExplainedByAI bool   `json:"explainedByAI"`
}

// RunCode writes the given Go source to a temp file, runs it with `go run`
// using the user's local Go toolchain, and captures the result. On failure,
// the raw compiler/runtime output is passed through the explain engine.
func (a *App) RunCode(source string) RunResult {
	tmpDir, err := os.MkdirTemp("", "gobug-run-*")
	if err != nil {
		return RunResult{Stderr: err.Error(), Success: false}
	}
	defer os.RemoveAll(tmpDir)

	mainFile := filepath.Join(tmpDir, "main.go")
	if err := os.WriteFile(mainFile, []byte(source), 0644); err != nil {
		return RunResult{Stderr: err.Error(), Success: false}
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "run", mainFile)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	runErr := cmd.Run()

	result := RunResult{
		Stdout:  stdout.String(),
		Stderr:  stderr.String(),
		Success: runErr == nil,
	}

	if runErr != nil {
		explanation, matched := engine.Explain(stderr.String())
		if !matched && a.cfg.APIKey != "" {
			aiExplanation, aiErr := engine.ExplainWithAI(a.cfg.APIKey, a.cfg.Model, source, stderr.String())
			if aiErr == nil {
				explanation = aiExplanation
				result.ExplainedByAI = true
			} else {
				explanation = explanation + "\n\nAI fallback failed: " + aiErr.Error()
			}
		} else if !matched {
			explanation = explanation + "\n\nAdd your own API key in Settings to get an AI explanation for cases like this."
		}
		result.Explanation = explanation
	}

	return result
}
