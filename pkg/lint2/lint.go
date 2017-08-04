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

type validator func() error

type linter interface {
	Load() error
	Lint() []error
}

// Lint lints the chart files that exist at the given path
func Lint(chartPath *string) []error {
	linters := []linter{
		newChartDir(chartPath),
		newChartFile(chartPath),
	}
	var violations []error
	for _, l := range linters {
		if loadErr := l.Load(); loadErr != nil {
			violations = append(violations, loadErr)
			// Could not load, skip linting
			continue
		}
		if lintErrs := l.Lint(); lintErrs != nil {
			violations = append(violations, lintErrs...)
		}
	}
	return violations
}
