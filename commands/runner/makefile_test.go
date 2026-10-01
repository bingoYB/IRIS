package runner

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/versenilvis/iris/spec"
)

func TestMakeTargetCompletion(t *testing.T) {
	cwd := t.TempDir()
	if err := os.WriteFile(filepath.Join(cwd, "Makefile"), []byte("gen-all:\n\t@echo done\n"), 0644); err != nil {
		t.Fatal(err)
	}
	spec.SetCWD(cwd)
	t.Cleanup(func() { spec.SetCWD("") })

	for _, input := range []string{"make ", "make gen"} {
		results := spec.Lookup(input)
		if len(results) != 1 || results[0].Cmd != "make gen-all" {
			t.Errorf("Lookup(%q) = %v, want only make gen-all", input, results)
		}
	}
}
