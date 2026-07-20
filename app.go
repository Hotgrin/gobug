package main

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"

	"gobug/engine"
)

// App holds the Wails runtime context and is bound to the frontend.
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// RunResult is sent back to the frontend after a run attempt.
type RunResult struct {
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
	Success     bool   `json:"success"`
	Explanation string `json:"explanation"`
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
		result.Explanation = engine.Explain(stderr.String())
	}

	return result
}
