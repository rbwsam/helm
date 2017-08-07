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

// The result of linting a single chart dir
type result struct {
	ChartDir        string      // Path to the chart dir
	HighestSeverity int         // Highest severity in the list of violations
	Violations      []violation // List of violations across various files
}

func newResult(path string) result {
	return result{ChartDir: path, HighestSeverity: 0, Violations: []violation{}}
}
