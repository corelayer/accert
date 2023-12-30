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

package config

import (
	"bytes"
	"log/slog"

	"github.com/corelayer/go-application/pkg/base"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	"github.com/corelayer/accert/cmd/accert/shared"
	"github.com/corelayer/accert/pkg/global"
)

var EncryptCommand = base.Command{
	Cobra: &cobra.Command{
		Use:           "encrypt",
		Short:         "Encrypt config",
		Long:          "ACME Protocol-based Certificate Manager - Encrypt config",
		RunE:          executeEncrypt,
		SilenceErrors: true,
		SilenceUsage:  true,
	},
	SubCommands: nil,
}

func executeEncrypt(cmd *cobra.Command, args []string) error {
	slog.Info("application started")
	defer slog.Info("application terminated")
	slog.Warn("flags", "set", shared.RootFlags)

	var (
		err     error
		hexData string

		globalViper *viper.Viper
		envViper    *viper.Viper
	)

	globalViper = base.NewConfiguration(shared.RootFlags.ConfigFile, shared.RootFlags.SearchFlag).GetViper()

	// Environment flags should go to different viper!!!!!
	envViper = viper.New()
	envViper.SetEnvPrefix(shared.APPLICATION_ENVIRONMENT_VARIABLE_PREFIX)
	err = envViper.BindEnv(shared.APPLICATION_ENVIRONMENT_ENCRYPTION_KEY)
	if err != nil {
		return err
	}

	err = globalViper.ReadInConfig()
	if err != nil {
		slog.Error("could not read config", "error", err.Error())
		return err
	}

	var config global.Configuration
	err = globalViper.Unmarshal(&config)
	if err != nil {
		return err
	}
	hexData, err = config.EncodeYamlToHex()
	if err != nil {
		return err
	}

	secureConfig := base.SecureData{
		Nonce:       "",
		CipherSuite: "AES_256_GCM",
		HexData:     hexData,
	}

	masterKey := envViper.GetString(shared.APPLICATION_ENVIRONMENT_ENCRYPTION_KEY)

	err = secureConfig.Encrypt(masterKey)
	if err != nil {
		return err
	}

	var secureData []byte
	secureData, err = yaml.Marshal(secureConfig)
	if err != nil {
		return err
	}

	err = globalViper.ReadConfig(bytes.NewBuffer(secureData))
	if err != nil {
		return err
	}

	return globalViper.WriteConfig()
	// return nil
}
