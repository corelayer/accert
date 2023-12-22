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
	"fmt"
	"log/slog"
	"os"

	"github.com/corelayer/go-application/pkg/base"
	"github.com/spf13/cobra"
)

var Command = base.Command{
	Cobra: &cobra.Command{
		Use:           "console",
		Short:         "Console mode",
		Long:          "ACME Protocol-based Certificate Manager - Console mode",
		RunE:          execute,
		SilenceErrors: true,
		SilenceUsage:  true,
	},
	SubCommands: nil,
}

func execute(cmd *cobra.Command, args []string) error {
	var (
		err    error
		logger *slog.Logger
	)
	fmt.Println("CONSOLE")

	logger, err = base.GetLogger(cmd, os.Stdout)
	if err != nil {
		return err
	}

	slog.SetDefault(logger)
	slog.Info("test")
	return nil
}
