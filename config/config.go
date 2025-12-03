package config

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Apikey     string `json:"APIKEY"`
	ConfigPath string
	HomeDir    string
}

func New() *Config {
	home, _ := os.UserHomeDir()
	return &Config{
		HomeDir:    home,
		ConfigPath: filepath.Join(home, ".config", "abuseipdb", "config.json"),
	}
}

func (c *Config) Load() {
	if _, err := os.Stat(c.ConfigPath); errors.Is(err, os.ErrNotExist) {
		println("configuration file not found! do a setup before running")
		println("exiting...")
		os.Exit(1)
	}
	raw, _ := os.ReadFile(c.ConfigPath)
	data := string(raw)
	json.NewDecoder(strings.NewReader(data)).Decode(c)
}

func (c *Config) Setup() {
	fmt.Print("Enter apikey: ")
	fmt.Scan(&c.Apikey)
}

func (c *Config) Save() {
	config_data, _ := json.MarshalIndent(c, "", "  ")
	user_readonly := os.FileMode(0700)
	dir, file := filepath.Split(c.ConfigPath)
	os.MkdirAll(dir, user_readonly)
	os.WriteFile(dir+"/"+file, config_data, user_readonly)
}
