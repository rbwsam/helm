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
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChartDir_Load(t *testing.T) {
	path := "testdata/albatross"
	cd := newChartDir(&path)
	err := cd.Load()
	assert.Nil(t, err)
}

func TestChartDir_Load_missing(t *testing.T) {
	path := "nowhere/fake"
	cd := newChartDir(&path)
	err := cd.Load()
	assert.Equal(t, fmt.Errorf("'%s' does not exist", path), err)
}

func TestChartDir_Load_notDir(t *testing.T) {
	path := "testdata/albatross/Chart.yaml"
	cd := newChartDir(&path)
	err := cd.Load()
	assert.Equal(t, fmt.Errorf("'%s' is not a directory", path), err)
}

func TestChartDir_Lint(t *testing.T) {
	path := "testdata/albatross"
	cd := newChartDir(&path)

	if err := cd.Load(); assert.Nil(t, err) {
		errs := cd.Lint()
		assert.Empty(t, errs)
	}
}
