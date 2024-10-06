package main

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Apps []string `toml:"apps"`
}

// LoadConfig 設定ファイル読み込み
func LoadConfig(path string) (*Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		config := defaultConfig()
		if err := config.Save(path); err != nil {
			return nil, err
		}
		return config, nil
	}

	var config *Config
	_, err = toml.DecodeFile(path, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Save 設定ファイル保存
func (c *Config) Save(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	err = toml.NewEncoder(file).Encode(c)
	if err != nil {
		return err
	}
	return nil
}

// defaultConfig 既定の設定ファイル
func defaultConfig() *Config {
	return &Config{Apps: []string{"notepad"}}
}
