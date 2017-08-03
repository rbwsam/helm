/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

import "fmt"

// Message describes an error encountered while linting.
type Message struct {
	// Severity is one of the *Sev constants
	Severity int
	Path     string
	Err      error
}

func (m Message) Error() string {
	return fmt.Sprintf("[%s] %s: %s", sev[m.Severity], m.Path, m.Err.Error())
}

// NewMessage creates a new Message struct
func NewMessage(severity int, path string, err error) Message {
	return Message{Severity: severity, Path: path, Err: err}
}
