package engine

import (
	"strings"
	"testing"
)

func TestExplainWithAI_NoKey(t *testing.T) {
	_, err := ExplainWithAI("", "", "package main", "some error")
	if err == nil {
		t.Fatal("expected an error when no API key is configured")
	}
}

func TestExtractContext_WithLineNumber(t *testing.T) {
	source := "line1\nline2\nline3\nline4\nline5\nline6\nline7"
	raw := "./main.go:4:2: some error"

	ctx := extractContext(source, raw)
	if !strings.Contains(ctx, "line4") {
		t.Fatalf("expected context to include the referenced line, got: %q", ctx)
	}
}

func TestExtractContext_NoLineNumber(t *testing.T) {
	source := "line1\nline2\nline3"
	raw := "runtime error: invalid memory address or nil pointer dereference"

	ctx := extractContext(source, raw)
	if ctx == "" {
		t.Fatal("expected a fallback context window even without a line number")
	}
}
