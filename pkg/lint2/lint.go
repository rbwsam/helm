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
	"k8s.io/helm/pkg/lint2/validations"
	"path"
)

// Severity indicates the severity of a Violation.
const (
	// UnknownSev indicates that the severity of the error is unknown, and should not stop processing.
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

type runnerFunc func()

type linterFunc func(path string) error

// All lints all of the relevant files in `chartPath`
func All(chartPath string) result {
	res := newResult(chartPath)
	chartFilePath := path.Join(chartPath, "Chart.yaml")

	//List of linters with paths and severity levels
	runners := []runnerFunc{
		runner(res, ErrorSev, chartFilePath, validations.NotDir),
	}

	for _, r := range runners {
		r()
	}
	return res
}

func runner(res result, sev int, path string, lin linterFunc) runnerFunc {
	return func() {
		if err := lin(path); err != nil {
			res.Violations = append(res.Violations, newViolation(sev, path, err))

			if sev > res.HighestSeverity {
				res.HighestSeverity = sev
			}
		}
	}
}
