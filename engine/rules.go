package engine

import "regexp"

// Rule matches a known Go compiler/vet/runtime message and turns it into a
// plain-English explanation. This is v0.1's offline tier — no API key needed.
type Rule struct {
	Name    string
	Pattern *regexp.Regexp
	Explain func(match []string) string
}

var rules = []Rule{
	{
		Name:    "undefined identifier",
		Pattern: regexp.MustCompile(`undefined: (\w+)`),
		Explain: func(m []string) string {
			return "Go can't find something named \"" + m[1] + "\". This usually means:\n" +
				"  - you misspelled it\n" +
				"  - you forgot to declare it (var/const/func) before using it\n" +
				"  - you meant to import a package that provides it, but didn't\n" +
				"Check spelling first — Go is case-sensitive, so `Foo` and `foo` are different names."
		},
	},
	{
		Name:    "type mismatch",
		Pattern: regexp.MustCompile(`cannot use (.+?) \(variable of type (\w+)\) as (\w+) value`),
		Explain: func(m []string) string {
			return "You're trying to use a value of type `" + m[2] + "` where Go expects a `" + m[3] + "`.\n" +
				"Go won't auto-convert types for you like some languages do. Fix: convert explicitly, e.g. " +
				"`" + m[3] + "(" + m[1] + ")`, or double-check you're assigning the right variable."
		},
	},
	{
		Name:    "unused import",
		Pattern: regexp.MustCompile(`"(.+?)" imported and not used`),
		Explain: func(m []string) string {
			return "You imported \"" + m[1] + "\" but never actually used it in the file.\n" +
				"Go treats unused imports as a compile error on purpose, to keep code clean. " +
				"Either use the package, or delete the import line."
		},
	},
	{
		Name:    "unused variable",
		Pattern: regexp.MustCompile(`declared and not used: (\w+)`),
		Explain: func(m []string) string {
			return "You declared a variable `" + m[1] + "` but never used it.\n" +
				"Go doesn't allow unused local variables. Either use it, or discard it with `_` " +
				"if you only need part of a return value: `_, err := someFunc()`."
		},
	},
	{
		Name:    "missing return",
		Pattern: regexp.MustCompile(`missing return`),
		Explain: func(m []string) string {
			return "This function promises a return type, but at least one path through it doesn't return a value.\n" +
				"Common cause: an `if` block returns, but there's no matching `else` or trailing return " +
				"for the remaining paths. Go checks every path at compile time — it won't guess for you."
		},
	},
	{
		Name:    "nil pointer panic",
		Pattern: regexp.MustCompile(`nil pointer dereference`),
		Explain: func(m []string) string {
			return "Runtime panic: you tried to use a pointer, map, or struct field that was never initialized (it's `nil`).\n" +
				"Common causes: a struct pointer declared with `var x *Type` but never set with `&Type{}`, " +
				"or reading a map key before the map was created with `make(map[...]...)`."
		},
	},
	{
		Name:    "index out of range",
		Pattern: regexp.MustCompile(`index out of range \[(\d+)\] with length (\d+)`),
		Explain: func(m []string) string {
			return "Runtime panic: you tried to access index " + m[1] + " of a slice/array that only has " + m[2] + " element(s).\n" +
				"Indexes start at 0, so a slice of length " + m[2] + " has valid indexes 0 through " + m[2] + "-1. " +
				"Double-check loop bounds — usually `i < len(x)`, not `i <= len(x)`."
		},
	},
}

// Explain matches raw compiler/vet/runtime output against known patterns and
// returns a plain-English explanation, or a generic fallback if nothing matches.
// The fallback is where a BYOK LLM call slots in for v0.2 — deliberately left
// as a single seam so it's a small change, not a rewrite.
func Explain(raw string) string {
	for _, r := range rules {
		if m := r.Pattern.FindStringSubmatch(raw); m != nil {
			return r.Explain(m)
		}
	}
	return "gobug doesn't recognize this one yet (v0.1 ships with a starter rule set).\n" +
		"The raw error is above. This is exactly the kind of case a BYOK AI explain step " +
		"(coming in v0.2) would take over for."
}
