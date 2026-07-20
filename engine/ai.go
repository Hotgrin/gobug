package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// DefaultAIModel is used when the user hasn't set one explicitly. Anthropic
// occasionally renames/retires model IDs — if this stops working, the user
// can override it in Settings. See https://docs.claude.com for current IDs.
const DefaultAIModel = "claude-sonnet-4-5"

var lineRefPattern = regexp.MustCompile(`:(\d+):\d+:`)

// ExplainWithAI sends the error plus a small window of surrounding source to
// the Anthropic API using the caller's own key, and returns a plain-English
// explanation. Deliberately sends only the relevant lines, not the whole file.
func ExplainWithAI(apiKey, model, source, rawError string) (string, error) {
	if apiKey == "" {
		return "", fmt.Errorf("no API key configured")
	}
	if model == "" {
		model = DefaultAIModel
	}

	context := extractContext(source, rawError)

	prompt := fmt.Sprintf(`A beginner Go developer got this error:

%s

Relevant code:
%s

Explain in plain English: (1) what went wrong, (2) why, (3) a concrete fix. Keep it under 120 words, no markdown headers, just plain text.`,
		strings.TrimSpace(rawError), context)

	reqBody, err := json.Marshal(map[string]interface{}{
		"model":      model,
		"max_tokens": 400,
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AI request failed (%d): %s", resp.StatusCode, string(respBody))
	}

	var parsed struct {
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	}
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return "", err
	}

	var out strings.Builder
	for _, block := range parsed.Content {
		if block.Type == "text" {
			out.WriteString(block.Text)
		}
	}
	if out.Len() == 0 {
		return "", fmt.Errorf("empty response from AI")
	}
	return out.String(), nil
}

// extractContext pulls a small window of source around the line referenced
// in the error (rather than sending the whole file). Falls back to a capped
// prefix of the source if no line number can be parsed from the error.
func extractContext(source, rawError string) string {
	lines := strings.Split(source, "\n")
	if m := lineRefPattern.FindStringSubmatch(rawError); m != nil {
		lineNum := 0
		fmt.Sscanf(m[1], "%d", &lineNum)
		start := lineNum - 3
		if start < 0 {
			start = 0
		}
		end := lineNum + 2
		if end > len(lines) {
			end = len(lines)
		}
		return strings.Join(lines[start:end], "\n")
	}
	max := 40
	if len(lines) < max {
		max = len(lines)
	}
	return strings.Join(lines[:max], "\n")
}
