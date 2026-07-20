package engine

import "testing"

func TestExplain_KnownPatterns(t *testing.T) {
	cases := []struct {
		name    string
		raw     string
		wantHit bool
	}{
		{"undefined identifier", `./main.go:6:6: undefined: fmt.Printlmn`, true},
		{"type mismatch", `cannot use x (variable of type int) as string value in assignment`, true},
		{"unused import", `"fmt" imported and not used`, true},
		{"unused variable", `declared and not used: x`, true},
		{"missing return", `missing return`, true},
		{"nil pointer panic", `runtime error: invalid memory address or nil pointer dereference`, true},
		{"index out of range", `runtime error: index out of range [5] with length 3`, true},
		{"unrecognized error", `some completely novel error message nobody wrote a rule for`, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			explanation, matched := Explain(tc.raw)
			if matched != tc.wantHit {
				t.Fatalf("Explain(%q) matched = %v, want %v", tc.raw, matched, tc.wantHit)
			}
			if explanation == "" {
				t.Fatalf("Explain(%q) returned empty explanation", tc.raw)
			}
		})
	}
}
