package js

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/versenilvis/iris/spec"
)

func TestPnpmDirectScriptCompletion(t *testing.T) {
	cwd := t.TempDir()
	spec.SetCWD(cwd)
	t.Cleanup(func() { spec.SetCWD("") })
	data := []byte(`{"scripts":{"dev":"vite","typecheck":"tsc --noEmit","lint:fix":"eslint . --fix"}}`)
	if err := os.WriteFile(filepath.Join(cwd, "package.json"), data, 0644); err != nil {
		t.Fatal(err)
	}

	for _, input := range []string{"pnpm ", "pnpm type", "pnpm lint:", "pnpm run "} {
		results := spec.Lookup(input)
		want := "pnpm typecheck"
		if input == "pnpm lint:" {
			want = "pnpm lint:fix"
		} else if input == "pnpm run " {
			want = "pnpm run typecheck"
		}
		if !hasSuggestion(results, want) {
			t.Errorf("Lookup(%q) did not suggest %q: %v", input, want, results)
		}
	}

	results := spec.Lookup("pnpm ")
	if !hasSuggestion(results, "pnpm install") || !hasSuggestion(results, "pnpm dev") {
		t.Errorf("built-in pnpm commands missing: %v", results)
	}
	if countSuggestion(results, "pnpm dev") != 1 {
		t.Errorf("pnpm dev appeared more than once: %v", results)
	}
}

func TestPnpmDirectScriptCompletionWithoutPackage(t *testing.T) {
	spec.SetCWD(t.TempDir())
	t.Cleanup(func() { spec.SetCWD("") })

	if hasSuggestion(spec.Lookup("pnpm lint"), "pnpm lint") {
		t.Error("direct script completion should require package.json")
	}
	if !hasSuggestion(spec.Lookup("pnpm run "), "pnpm run lint") {
		t.Error("pnpm run fallback suggestions changed")
	}
}

func hasSuggestion(results []spec.Suggestion, cmd string) bool {
	return countSuggestion(results, cmd) > 0
}

func countSuggestion(results []spec.Suggestion, cmd string) int {
	count := 0
	for _, result := range results {
		if result.Cmd == cmd {
			count++
		}
	}
	return count
}
