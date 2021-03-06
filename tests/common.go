// Copyright 2015 The appc Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tests

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	acbuildBinPath string
)

func init() {
	acbuildBinPath = os.Getenv("ACBUILD_BIN")
	if acbuildBinPath == "" {
		fmt.Fprintf(os.Stderr, "ACBUILD_BIN environmment variable must be set\n")
		os.Exit(1)
	} else if _, err := os.Stat(acbuildBinPath); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func runAcbuild(workingDir string, args ...string) error {
	cmd := exec.Command(acbuildBinPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = workingDir
	return cmd.Run()
}

func setUpTest() string {
	tmpdir := mustTempDir()

	err := runAcbuild(tmpdir, "begin")
	if err != nil {
		panic(err)
	}

	return tmpdir
}

func cleanUpTest(tmpdir string) error {
	return os.RemoveAll(tmpdir)
}

func mustTempDir() string {
	dir, err := ioutil.TempDir("", "acbuild-test")
	if err != nil {
		panic(err)
	}
	return dir
}
