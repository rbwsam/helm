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
	"errors"
	"fmt"
	"k8s.io/helm/pkg/chartutil"
	"os"
)

type valuesFile struct {
	path string
}

func newValuesFile(path string) *valuesFile {
	return &valuesFile{path}
}

func (vf *valuesFile) Load() error {
	loaders := []loader{
		vf.checkNotDir,
		vf.parse,
	}
	for _, l := range loaders {
		if err := l(); err != nil {
			return err
		}
	}
	return nil
}

func (vf *valuesFile) Lint() []Violation {
	return []Violation{}
}

func (vf *valuesFile) HighestSeverity() int {
	return UnknownSev
}

func (vf *valuesFile) checkNotDir() error {
	fi, err := os.Stat(vf.path)

	if err == nil && fi.IsDir() {
		return errors.New("should be a file, not a directory")
	}
	return nil
}

func (vf *valuesFile) parse() error {
	_, err := chartutil.ReadValuesFile(vf.path)
	if err != nil {
		return fmt.Errorf("unable to parse YAML\n\t%s", err)
	}
	return nil
}
