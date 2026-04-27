// Copyright 2022 Google LLC
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

// osv-scanner checks projects and dependencies for known vulnerabilities
// using the OSV (Open Source Vulnerabilities) database.
package main

import (
	"os"

	"github.com/google/osv-scanner/cmd/osv-scanner/internal/cmd"
)

func main() {
	if err := cmd.Run(os.Args, os.Stdout, os.Stderr); err != nil {
		os.Exit(1)
	}
}
