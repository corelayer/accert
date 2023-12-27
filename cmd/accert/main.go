/*
 * Copyright 2023 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package main

import (
	"os"
	"path/filepath"

	"github.com/corelayer/go-application/pkg/base"

	"github.com/corelayer/accert/cmd/accert/cmd/console"
	"github.com/corelayer/accert/cmd/accert/cmd/daemon"
)

const (
	APPLICATION_NAME   = "accert"
	APPLICATION_TITLE  = "ACME Protocol-based Certificate Manager"
	APPLICATION_BANNER = "ACME Protocol-based Certificate Manager"
)

var (
	configSearchPaths = []string{
		filepath.Join("$HOME", ".lens"),
		filepath.Join("$PWD"),
	}
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	var (
		// err error

		configFileFlag   string
		configSearchFlag []string
	)
	app := base.NewApplication(APPLICATION_NAME, APPLICATION_TITLE, APPLICATION_BANNER, "")
	app.Command.PersistentFlags().StringVarP(&configFileFlag, "configFile", "", "config.yaml", "config file name")
	app.Command.PersistentFlags().StringSliceVarP(&configSearchFlag, "configSearchPath", "", configSearchPaths, "config file search paths")

	app.RegisterCommands([]base.Commander{
		daemon.Command,
		console.Command,
	})
	return app.Run()
}
