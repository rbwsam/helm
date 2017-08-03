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

package support

// Linter encapsulates a linting run of a particular chart.
type Linter struct {
	Messages []Message
	// The highest severity of all the failing lint rules
	HighestSeverity int
	ChartDir        string
}

// RunLinterRule returns true if the validation passed
func (l *Linter) RunLinterRule(severity int, path string, err error) bool {
	// severity is out of bound
	if severity < 0 || severity >= len(sev) {
		return false
	}

	if err != nil {
		l.Messages = append(l.Messages, NewMessage(severity, path, err))

		if severity > l.HighestSeverity {
			l.HighestSeverity = severity
		}
	}
	return err == nil
}
