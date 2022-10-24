// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
)

const requiredF = `required flag(s) "%s" not set`

var errMissingProjectID = fmt.Errorf(requiredF, flag.ProjectID)
var ErrMissingOrgID = fmt.Errorf(requiredF, flag.OrgID)
var ErrFreeClusterAlreadyExists = errors.New("this project already has another free cluster")
var ErrNoRegionExistsTryCommand = errors.New(`the region does not exist. to find the available regions, run "atlas cluster availableRegions list --help"`)
