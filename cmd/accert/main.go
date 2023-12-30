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

	"github.com/corelayer/accert/cmd/accert/cmd/config"
	"github.com/corelayer/accert/cmd/accert/cmd/console"
	"github.com/corelayer/accert/cmd/accert/cmd/daemon"
	"github.com/corelayer/accert/cmd/accert/shared"
)

var (
	configSearchPaths = []string{
		filepath.Join("$HOME", ".accert"),
		filepath.Join("$PWD"),
	}
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	app := base.NewApplication(shared.APPLICATION_NAME, shared.APPLICATION_TITLE, shared.APPLICATION_BANNER, "")
	app.Command.PersistentFlags().StringVarP(&shared.RootFlags.ConfigFile, "configFile", "", "config.yaml", "config file name")
	app.Command.PersistentFlags().StringSliceVarP(&shared.RootFlags.SearchFlag, "configSearchPath", "", configSearchPaths, "config file search paths")
	// app.Command.PersistentFlags().BoolVarP(&shared.RootFlags.Encrypted, "encrypted", "e", false, "config file encrypted")

	app.RegisterCommands([]base.Commander{
		daemon.Command,
		console.Command,
		config.Command,
	})
	return app.Run()
}
