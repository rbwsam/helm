package lint2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChartFile_Load(t *testing.T) {
	cf := newChartFile("testdata/albatross/Chart.yaml")
	err := cf.Load()
	assert.Nil(t, err)
}

func TestChartFile_Load_isDir(t *testing.T) {
	cf := newChartFile("testdata/albatross")
	err := cf.Load()
	assert.EqualError(t, err, "should be a file, not a directory")
}

func TestChartFile_Lint(t *testing.T) {
	cf := newChartFile("testdata/albatross/Chart.yaml")
	if assert.Nil(t, cf.Load()) {
		cf.Lint()
		assert.Empty(t, cf.Violations)
	}
}
