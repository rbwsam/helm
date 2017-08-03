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

import (
	"errors"
	"testing"
)

func TestMessage(t *testing.T) {
	m := Message{ErrorSev, "Chart.yaml", errors.New("Foo")}
	if m.Error() != "[ERROR] Chart.yaml: Foo" {
		t.Errorf("Unexpected output: %s", m.Error())
	}

	m = Message{WarningSev, "templates/", errors.New("Bar")}
	if m.Error() != "[WARNING] templates/: Bar" {
		t.Errorf("Unexpected output: %s", m.Error())
	}

	m = Message{InfoSev, "templates/rc.yaml", errors.New("FooBar")}
	if m.Error() != "[INFO] templates/rc.yaml: FooBar" {
		t.Errorf("Unexpected output: %s", m.Error())
	}
}
