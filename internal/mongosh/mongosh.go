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
package mongosh

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

const (
	mongoshBin = "mongosh"
)

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func Bin() string {
	if isWindows() {
		return fmt.Sprintf("%s.exe", mongoshBin)
	}
	return mongoshBin
}

func Path() string {
	if path, err := exec.LookPath(Bin()); err == nil {
		return path
	}

	return ""
}

func Run(binary, username, password, mongoURI string) error {
	args := []string{"mongosh", "-u", username, "-p", password, mongoURI}
	env := os.Environ()
	return syscall.Exec(binary, args, env)
}
