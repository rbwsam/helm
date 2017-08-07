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
	"github.com/Masterminds/semver"
	"github.com/asaskevich/govalidator"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type chartFile struct {
	Violations      []violation
	HighestSeverity int
	path            string
	metadata        *chart.Metadata
}

func newChartFile(path string) *chartFile {
	return &chartFile{path: path, HighestSeverity: UnknownSev}
}

func (cf *chartFile) Load() error {
	if err := cf.checkNotDir(); err != nil {
		return err
	}
	if err := cf.parse(); err != nil {
		return err
	}
	return nil
}

func (cf *chartFile) Lint() {
	scoredLinters := []scoredLinter{
		newScoredLinter(ErrorSev, cf.lintName),
		newScoredLinter(ErrorSev, cf.lintDirName),
		newScoredLinter(ErrorSev, cf.lintVersion),
		newScoredLinter(ErrorSev, cf.lintEngine),
		newScoredLinter(ErrorSev, cf.lintMaintainers),
		newScoredLinter(ErrorSev, cf.lintSources),
		newScoredLinter(InfoSev, cf.lintIconPresence),
		newScoredLinter(ErrorSev, cf.lintIconURL),
	}

	for _, sc := range scoredLinters {
		if err := sc.linter(); err != nil {
			v := newViolation(sc.severity, cf.path, err)
			cf.Violations = append(cf.Violations, v)
		}
	}
}

func (cf *chartFile) checkNotDir() error {
	fi, err := os.Stat(cf.path)

	if err == nil && fi.IsDir() {
		return errors.New("should be a file, not a directory")
	}
	return nil
}

func (cf *chartFile) parse() error {
	metadata, err := chartutil.LoadChartfile(cf.path)
	if err != nil {
		return err
	}
	cf.metadata = metadata
	return nil
}

func (cf *chartFile) lintName() error {
	if cf.metadata.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (cf *chartFile) lintDirName() error {
	dirName := filepath.Base(path.Join(cf.path, "../"))
	if cf.metadata.Name != dirName {
		return fmt.Errorf("directory name (%s) and chart name (%s) must be the same", dirName, cf.metadata.Name)
	}
	return nil
}

func (cf *chartFile) lintVersion() error {
	if cf.metadata.Version == "" {
		return errors.New("version is required")
	}

	version, err := semver.NewVersion(cf.metadata.Version)

	if err != nil {
		return fmt.Errorf("version '%s' is not a valid SemVer", cf.metadata.Version)
	}

	c, err := semver.NewConstraint("> 0")
	if err != nil {
		return err
	}
	valid, msg := c.Validate(version)

	if !valid && len(msg) > 0 {
		return fmt.Errorf("version %v", msg[0])
	}

	return nil
}

func (cf *chartFile) lintEngine() error {
	if cf.metadata.Engine == "" {
		return nil
	}

	keys := make([]string, 0, len(chart.Metadata_Engine_value))
	for engine := range chart.Metadata_Engine_value {
		str := strings.ToLower(engine)

		if str == "unknown" {
			continue
		}

		if str == cf.metadata.Engine {
			return nil
		}

		keys = append(keys, str)
	}

	return fmt.Errorf("engine '%v' not valid. Valid options are %v", cf.metadata.Engine, keys)
}

func (cf *chartFile) lintMaintainers() error {
	for _, maintainer := range cf.metadata.Maintainers {
		if maintainer.Name == "" {
			return errors.New("each maintainer requires a name")
		} else if maintainer.Email != "" && !govalidator.IsEmail(maintainer.Email) {
			return fmt.Errorf("invalid email '%s' for maintainer '%s'", maintainer.Email, maintainer.Name)
		}
	}
	return nil
}

func (cf *chartFile) lintSources() error {
	for _, source := range cf.metadata.Sources {
		if source == "" || !govalidator.IsRequestURL(source) {
			return fmt.Errorf("invalid source URL '%s'", source)
		}
	}
	return nil
}

func (cf *chartFile) lintIconPresence() error {
	if cf.metadata.Icon == "" {
		return errors.New("icon is recommended")
	}
	return nil
}

func (cf *chartFile) lintIconURL() error {
	if cf.metadata.Icon != "" && !govalidator.IsRequestURL(cf.metadata.Icon) {
		return fmt.Errorf("invalid icon URL '%s'", cf.metadata.Icon)
	}
	return nil
}
