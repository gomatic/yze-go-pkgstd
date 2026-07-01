package pkgstd_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"

	pkgstd "github.com/gomatic/yze-go-pkgstd"
)

func TestCommandPackageStandards(t *testing.T) {
	analysistest.Run(
		t, analysistest.TestData(), pkgstd.Analyzer,
		"m/internal/domain/greet",
		"m/internal/app/commands/greet",
		"m/internal/app/commands/multifile",
		"m/internal/app/commands/multiverb",
		"m/internal/app/commands/badalias",
		"m/internal/app/commands/noconst",
		"m/internal/app/commands/nocmd",
		"m/internal/app/commands/examples",
	)
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, pkgstd.Registration.Validate())
	assert.Equal(t, "yze/pkgstd", pkgstd.Registration.RuleID())
	assert.Same(t, pkgstd.Analyzer, pkgstd.Registration.Analyzer)
}
