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
	"path"
)

// Severity indicates the Severity of a Message.
const (
	// UnknownSev indicates that the Severity of the error is unknown, and should not stop processing.
	UnknownSev = iota
	// InfoSev indicates information, for example missing values.yaml file
	InfoSev
	// WarningSev indicates that something does not meet code standards, but will likely function.
	WarningSev
	// ErrorSev indicates that something will not likely function.
	ErrorSev
)

// sev matches the *Sev states.
var sev = []string{"UNKNOWN", "INFO", "WARNING", "ERROR"}

type lintable interface {
	Load() error
	Lint() []Violation
	HighestSeverity() int
}

type linter func() error
type loader func() error

func Lint(chartDir string) (Result, error) {
	result := newResult()

	lintables := []lintable{
		newChartFile(path.Join(chartDir, "Chart.yaml")),
		newValuesFile(path.Join(chartDir, "values.yaml")),
	}

	for _, lintable := range lintables {
		if err := lintable.Load(); err != nil {
			return result, err
		}
		violations := lintable.Lint()
		result.Violations = append(result.Violations, violations...)
		if lintable.HighestSeverity() > result.HighestSeverity {
			result.HighestSeverity = lintable.HighestSeverity()
		}
	}
	return result, nil
}
