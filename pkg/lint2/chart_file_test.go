/*
Copyright 2017 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

func TestChartFile_Load_missing(t *testing.T) {
	cf := newChartFile("testdata/albatross/asdf.yaml")
	err := cf.Load()
	assert.EqualError(t, err, "open testdata/albatross/asdf.yaml: no such file or directory")
}

func TestChartFile_Load_isDir(t *testing.T) {
	cf := newChartFile("testdata/albatross")
	err := cf.Load()
	assert.EqualError(t, err, "should be a file, not a directory")
}

func TestChartFile_Load_isNotYaml(t *testing.T) {
	cf := newChartFile("testdata/albatross/templates/_helpers.tpl")
	err := cf.Load()
	assert.EqualError(t, err, "error converting YAML to JSON: yaml: did not find expected alphabetic or numeric character")
}

func TestChartFile_Lint(t *testing.T) {
	cf := newChartFile("testdata/albatross/Chart.yaml")
	if assert.Nil(t, cf.Load()) {
		violations := cf.Lint()
		assert.Empty(t, violations)
	}
}
