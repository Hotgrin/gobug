# Contributing to gobug

Thanks for considering it — the fastest way to make gobug better is to
report a real error message it doesn't explain well yet.

## Reporting an error gobug got wrong or missed

Open an issue with:
- The raw compiler/vet/runtime error text
- What you expected the explanation to say
- Go version (`go version`)

This is genuinely the highest-value contribution right now — the offline
rule set is only as good as the real-world errors it's been tested against.

## Adding a new rule

Rules live in `engine/rules.go`. Each one is a regex pattern matched
against raw `go build`/`go vet`/runtime output, plus a function that turns
the match into a plain-English explanation:

```go
{
    Name:    "short description",
    Pattern: regexp.MustCompile(`your regex here`),
    Explain: func(m []string) string {
        return "Plain-English explanation, referencing m[1], m[2], etc. " +
            "for captured groups from the regex."
    },
},
```

Guidelines for a good rule:
- Explain the **cause**, not just restate the error
- Give a **concrete fix**, not just "check your code"
- Keep it short — a beginner should be able to read it in one glance
- Add a test case in `engine/rules_test.go` covering the exact error string
  Go actually produces (not an approximation of it)

## Running tests locally

```
go vet ./...
go test ./... -v
gofmt -l .
```

All three run in CI on every PR.

## Code style

Standard `gofmt`. No linter config beyond that yet — keeping the bar low
on purpose so contributing a rule stays a five-minute job.
