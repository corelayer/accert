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

package global

import (
	"encoding/hex"
	"fmt"

	"gopkg.in/yaml.v3"
)

var TemplateConfiguration Configuration

type Configuration struct {
	Name    string   `json:"name" yaml:"name" mapstructure:"name"`
	Tenants []Tenant `json:"tenants" yaml:"tenants" mapstructure:"tenants"`
}

func (c Configuration) MarshalToYaml() ([]byte, error) {
	var (
		err  error
		data []byte
	)
	data, err = yaml.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("could not marshal configuration to yaml: %w", err)
	}
	return data, nil
}

func (c Configuration) EncodeYamlToHex() (string, error) {
	var (
		err  error
		data []byte
	)
	data, err = c.MarshalToYaml()
	if err != nil {
		return "", fmt.Errorf("could not encode yaml to hex: %w", err)
	}
	return hex.EncodeToString(data), nil
}

func NewConfigurationFromHex(s string) (Configuration, error) {
	var (
		err  error
		data []byte
		c    Configuration
	)

	data, err = hex.DecodeString(s)
	if err != nil {
		return Configuration{}, fmt.Errorf("could not decode configuration from hex: %w", err)
	}

	err = yaml.Unmarshal(data, c)
	if err != nil {
		return Configuration{}, fmt.Errorf("could not marshal data into configuration: %w", err)
	}
	return c, nil
}

func init() {
	TemplateConfiguration = Configuration{
		Name: "template",
		Tenants: []Tenant{
			{
				Name: "template",
				Environments: []Environment{
					{
						Name: "template",
						Connection: Connection{
							Name:        "template",
							Credentials: "template",
						},
					},
				},
				Credentials: []Credentials{
					{
						Name:     "template",
						Username: "username",
						Password: "password",
					},
				},
			},
		},
	}
}
