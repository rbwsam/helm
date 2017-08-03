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
)

// ChartDir encapsulates a linting run of a particular chart directory.
type ChartDir struct {
	path string
}

// NewChartDir returns a new ChartDir with the given path.
func NewChartDir(path string) (*ChartDir, error) {
	exists, err := dirExists(path)
	if err != nil {
		return &ChartDir{}, err
	}
	if exists {
		return &ChartDir{path}, nil
	}
	return &ChartDir{}, fmt.Errorf("'%s' does not exist", path)
}
