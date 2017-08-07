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
	"path/filepath"
	"strings"
)

type chartFile struct {
	path     *string
	metadata *chart.Metadata
}

func newChartFile(path *string) *chartFile {
	return &chartFile{path: path}
}

// Load loads the contents of the chartFile
func (cf *chartFile) Load() error {
	if err := cf.exists(); err != nil {
		return err
	}
	metadata, err := chartutil.LoadChartfile(*cf.path)
	if err != nil {
		return err
	}
	cf.metadata = metadata
	return nil
}

// Lint lints the chartFile's metadata via several linters
func (cf *chartFile) Lint() []error {
	validators := []validator{
		cf.validateName,
		cf.validateNameDirMatch,
		cf.validateVersion,
		cf.validateEngine,
		cf.validateMaintainers,
		cf.validateSources,
		cf.validateIcon,
	}

	var violations []error
	for _, v := range validators {
		if err := v(); err != nil {
			violations = append(violations, err)
		}
	}
	return violations
}

func (cf *chartFile) exists() error {
	exists, err := fileExists(*cf.path)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("'%s' does not exist", *cf.path)
	}
	return nil
}

func (cf *chartFile) validateName() error {
	if cf.metadata.Name == "" {
		return errors.New("name is required")
	}
	return nil
}

func (cf *chartFile) validateNameDirMatch() error {
	dirName := filepath.Base(filepath.Join(*cf.path, "../"))
	if cf.metadata.Name != dirName {
		return fmt.Errorf("directory name (%s) and chart name (%s) must be the same", dirName, cf.metadata.Name)
	}
	return nil
}

func (cf *chartFile) validateVersion() error {
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

func (cf *chartFile) validateEngine() error {
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

func (cf *chartFile) validateMaintainers() error {
	for _, maintainer := range cf.metadata.Maintainers {
		if maintainer.Name == "" {
			return errors.New("each maintainer requires a name")
		} else if maintainer.Email != "" && !govalidator.IsEmail(maintainer.Email) {
			return fmt.Errorf("invalid email '%s' for maintainer '%s'", maintainer.Email, maintainer.Name)
		}
	}
	return nil
}

func (cf *chartFile) validateSources() error {
	for _, source := range cf.metadata.Sources {
		if source == "" || !govalidator.IsRequestURL(source) {
			return fmt.Errorf("invalid source URL '%s'", source)
		}
	}
	return nil
}

func (cf *chartFile) validateIcon() error {
	if cf.metadata.Icon == "" {
		return errors.New("icon is recommended")
	}
	if !govalidator.IsRequestURL(cf.metadata.Icon) {
		return fmt.Errorf("invalid icon URL '%s'", cf.metadata.Icon)
	}
	return nil
}
