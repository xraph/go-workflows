// Package fixture proves the workflow can provision a JS toolchain for Go
// tests that shell out to it. A Makefile is present but deliberately defines
// a `test` target that would FAIL, so a run that passes also proves
// prefer-makefile:false genuinely bypassed it.
package fixture

import (
	"os/exec"
	"strings"
	"testing"
)

func TestNodeIsOnPath(t *testing.T) {
	out, err := exec.Command("node", "--version").Output()
	if err != nil {
		t.Fatalf("node not available: %v", err)
	}
	if !strings.HasPrefix(string(out), "v") {
		t.Fatalf("unexpected node version output: %q", out)
	}
}

func TestGlobalNpmPackageInstalled(t *testing.T) {
	out, err := exec.Command("esbuild", "--version").Output()
	if err != nil {
		t.Fatalf("esbuild not available; npm-global-packages did not install it: %v", err)
	}
	if len(strings.TrimSpace(string(out))) == 0 {
		t.Fatal("esbuild reported an empty version")
	}
}
