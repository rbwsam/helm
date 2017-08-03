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

func TestNewChartDir(t *testing.T) {
	path := "testdata/albatross"
	cd, err := NewChartDir(path)
	assert.Equal(t, path, cd.path)
	assert.Nil(t, err)
}

func TestNewChartDir_NotExist(t *testing.T) {
	path := "somewhere/fake"
	_, err := NewChartDir(path)
	expected := fmt.Errorf("'%s' does not exist", path)
	assert.Equal(t, expected, err)
}

func TestNewChartDir_NotDir(t *testing.T) {
	path := "testdata/albatross/Chart.yaml"
	_, err := NewChartDir(path)
	expected := fmt.Errorf("'%s' is not a directory", path)
	assert.Equal(t, expected, err)
}
