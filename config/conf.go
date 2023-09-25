// Copyright 2023 The Casdoor Authors. All Rights Reserved.
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

package config

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Endpoint     string `yaml:"endpoint"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Organization string `yaml:"organization"`
	Application  string `yaml:"application"`
	FrontendURL  string `yaml:"frontend_url"`
}

type Config struct {
	Certificate string       `yaml:"certificate"`
	Server      ServerConfig `yaml:"server"`
}

var GlobalConfig *Config

func LoadConfig(configPath string) error {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	GlobalConfig = &cfg

	return nil
}
