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

package console

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log/slog"

	"github.com/corelayer/go-application/pkg/base"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"

	"github.com/corelayer/accert/cmd/accert/shared"
)

var Command = base.Command{
	Cobra: &cobra.Command{
		Use:           "console",
		Short:         "Console mode",
		Long:          "ACME Protocol-based Certificate Manager - Console mode",
		RunE:          execute,
		SilenceErrors: true,
		SilenceUsage:  true,
		Annotations: map[string]string{
			"logtarget": "console.log",
		},
	},
	SubCommands: nil,
	Configure:   configureConsole,
}

func execute(cmd *cobra.Command, args []string) error {
	slog.Info("application started")
	defer slog.Info("application terminated")
	slog.Warn("flags", "set", shared.RootFlags)

	rootConfig := base.NewConfiguration(shared.RootFlags.ConfigFile, shared.RootFlags.SearchFlag)
	rootViper := rootConfig.GetViper()

	// Environment flags should go to different viper!!!!!
	envViper := viper.New()
	envViper.SetEnvPrefix("accert")
	envViper.BindEnv("key")

	var err error
	err = rootViper.ReadInConfig()
	if err != nil {
		slog.Error("could not read config", "error", err.Error())
		return err
	}

	var rootData base.SecureData
	err = rootViper.Unmarshal(&rootData)
	if err != nil {
		slog.Error("could not unmarshal config", "error", err.Error())
		return err
	}

	fmt.Println(rootData)

	masterKey := envViper.GetString("key")
	fmt.Println(masterKey)
	var decodedMaster []byte
	decodedMaster, err = hex.DecodeString(masterKey)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(decodedMaster)
	err = rootData.Decrypt(masterKey)
	// err = rootData.Encrypt(masterKey)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	fmt.Println(rootData)

	var rootBytes []byte
	rootBytes, err = yaml.Marshal(rootData)
	err = rootViper.ReadConfig(bytes.NewBuffer(rootBytes))
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = rootViper.WriteConfig()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func configureConsole(cmd *cobra.Command) {
	base.AddLogTargetFlag(cmd)
}
