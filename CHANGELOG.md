# Changelog

## v0.2

- Added BYOK AI fallback: when the offline rule set doesn't recognize an
  error, gobug can call the Anthropic API with your own key for a tailored
  explanation. Only the error and a small window of surrounding source are
  sent, never the whole file.
- Added Settings modal for managing the API key and model locally
  (stored in the OS config dir, never in the repo).
- Added unit tests for the rule engine, the AI-fallback context extraction,
  and config load/save.
- CI now runs `go test ./...` in addition to `go vet` and gofmt checks.

## v0.1

- Initial release: native "run area" built with Wails — write/paste Go,
  run it against your local toolchain, see output.
- Offline rule-based explain engine for 7 common beginner errors:
  undefined identifiers, type mismatches, unused imports/variables,
  missing return, nil pointer panics, index-out-of-range panics.
- Line-numbered editor, copy-error button.
- MIT licensed, CI on push (vet + gofmt).
