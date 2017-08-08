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
	"k8s.io/helm/pkg/proto/hapi/chart"
	"os"
)

type templateDir struct {
	path     string
	metadata *chart.Metadata
}

func newTemplateDir(path string) *templateDir {
	return &templateDir{path: path}
}

func (td *templateDir) Load() error {
	return nil
}

func (td *templateDir) Lint() ([]Violation, int) {
	scoredLinters := []scoredLinter{
		newScoredLinter(WarningSev, td.lintDir),
	}
	violations := []Violation{}
	highestSeverity := UnknownSev

	for _, sc := range scoredLinters {
		if err := sc.Linter(); err != nil {
			v := newViolation(sc.Severity, td.path, err)
			violations = append(violations, v)
			if sc.Severity > highestSeverity {
				highestSeverity = sc.Severity
			}
		}
	}
	return violations, highestSeverity
}

func (td *templateDir) lintDir() error {
	if fi, err := os.Stat(td.path); err != nil {
		return errors.New("directory not found")
	} else if err == nil && !fi.IsDir() {
		return errors.New("not a directory")
	}
	return nil
}
