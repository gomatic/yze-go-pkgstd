package nocmd

import "testing"

// TestCommand produces the in-package test variant (command.go plus a _test.go
// file) and matches the exported *Command entry-point shape. The analyzer must
// skip _test.go files when locating the command file, so this test function is
// never mistaken for the entry point: the variant pass reports the same
// missing-Command diagnostic as the primary pass — anchored at command.go — and
// the driver collapses the two onto the single `want` there.
func TestCommand(t *testing.T) { _ = t }
