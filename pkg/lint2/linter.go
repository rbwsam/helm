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

// Linter lints the chart files that exist at the given path
func Linter(chartPath *string) ([]string, error) {
	var violations []string
	linters := []linter{
		newChartDir(chartPath),
	}

	for _, l := range linters {
		v, err := l.Lint()
		// If we encounter an error, bail out immediately
		if err != nil {
			return violations, err
		}
		// If we get violations, add them to the list
		if len(v) > 0 {
			violations = append(violations, v...)
		}
	}
	return violations, nil
}
