# gobug

[![CI](https://github.com/hotgrin/gobug/actions/workflows/ci.yml/badge.svg)](https://github.com/hotgrin/gobug/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

The friendliest Go error explainer around. Run your code, and instead of a
raw compiler dump or a bare stack trace, get a plain-English explanation and
a fix — right next to your code.

v0.1 is a native desktop "run area": paste or write Go on the left, hit Run,
see output (and, on failure, a real explanation) on the right. No hosted
sandbox, no untrusted code execution risk — it shells out to *your own*
local `go` toolchain, same as typing `go run` yourself.

## Why this exists

Go already has excellent tooling for *finding* problems: `go build`,
`go vet`, `golangci-lint`. What's missing is a good layer for *explaining*
them to someone who doesn't yet read Go compiler output fluently. gobug is
that layer — not a rival static analyzer, a translator on top of the tools
that already exist.

## v0.1 scope (intentionally small)

- Single-file "run area" — write/paste Go, click Run
- Offline rule-based explain engine covering common beginner errors:
  undefined identifiers, type mismatches, unused imports/variables, missing
  return, nil pointer panics, index-out-of-range panics
- No API key required to be useful

## Not yet built (roadmap, not promises)

- BYOK LLM fallback for errors the rule set doesn't recognize yet
  (the seam is already there — see the fallback branch in
  `engine/rules.go`'s `Explain` function)
- Piping `golangci-lint` output through the same explain layer
- Proper code editor (syntax highlighting) instead of a plain textarea
- Multi-file / whole-package support instead of single-file snippets

## Building it

Clone it first:

```
git clone https://github.com/hotgrin/gobug.git
cd gobug
```

Prerequisites: Go 1.22+, Node.js (for Wails' internal tooling, even though
this frontend has no npm build step), and the Wails CLI:

```
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

Then, from this directory:

```
go mod tidy       # resolves the exact Wails version for your platform
wails dev         # live dev mode
wails build        # produces a standalone binary in build/bin
                    # (.exe on Windows, no extension on Linux/Mac)
```

On Windows, Wails apps need WebView2 (pre-installed on Windows 10/11 in
almost all cases — Wails will prompt if it's missing).

## Distribution plan (once v0.1 feels solid)

Same playbook as `goscan`: GitHub release with cross-platform binaries,
submit to Awesome-Go, post to dev.to and Show HN under #golang, drop it in
Gopher Slack's #showcase. Open source, BYOK where AI is involved, MIT
license.
