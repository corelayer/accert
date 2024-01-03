/*
 * Copyright 2024 CoreLayer BV
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

package config

import (
	"bytes"
	"log/slog"
	"os"

	"github.com/corelayer/go-application/pkg/base"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/corelayer/accert/cmd/accert/shared"
	"github.com/corelayer/accert/pkg/global"
)

var GlobalCreateCommand = base.Command{
	Cobra: &cobra.Command{
		Use:           "create",
		Short:         "create config",
		Long:          "ACME Protocol-based Certificate Manager - Global create config",
		RunE:          executeGlobalCreate,
		SilenceErrors: true,
		SilenceUsage:  true,
	},
	SubCommands: nil,
}

func executeGlobalCreate(cmd *cobra.Command, args []string) error {
	slog.Info("application started")
	defer slog.Info("application terminated")

	var (
		err         error
		globalViper *viper.Viper
		envViper    *viper.Viper
		data        []byte
	)
	_, err = os.Create(shared.RootFlags.ConfigFile)
	if err != nil {
		return err
	}

	globalViper = base.NewConfiguration(shared.RootFlags.ConfigFile, shared.RootFlags.SearchFlag).GetViper()

	// Environment flags should go to different viper!!!!!
	envViper = viper.New()
	envViper.SetEnvPrefix(shared.APPLICATION_ENVIRONMENT_VARIABLE_PREFIX)
	err = envViper.BindEnv(shared.APPLICATION_ENVIRONMENT_ENCRYPTION_KEY)
	if err != nil {
		return err
	}

	data, err = global.TemplateConfiguration.MarshalToYaml()
	if err != nil {
		return err
	}

	err = globalViper.ReadConfig(bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return globalViper.WriteConfig()
}
