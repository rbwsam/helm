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

import "fmt"

type chartDir struct {
	path *string
}

func newChartDir(path *string) *chartDir {
	return &chartDir{path}
}

func (cd *chartDir) Load() error {
	return cd.exists()
}

// Lint lints the ChartDir
func (cd *chartDir) Lint() []error {
	// No linting necessary for a directory
	return []error{}
}

func (cd *chartDir) exists() error {
	exists, err := dirExists(*cd.path)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("'%s' does not exist", *cd.path)
	}
	return nil
}
