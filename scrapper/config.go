package scrapper

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const configDir = "invest-scrapper-config"

type Config struct {
	Dir string
}

func InitConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dir := fmt.Sprintf("%s/%s", home, configDir)
	if err = os.Mkdir(dir, os.ModePerm); !errors.Is(err, os.ErrExist) && err != nil {
		return nil, err
	}
	return &Config{Dir: dir}, nil
}

func (c *Config) FilePath(filename string) string {
	return fmt.Sprintf("%s/%s", c.Dir, filename)
}

func GetWallet(walletPath string) ([]string, error) {
	content, err := os.ReadFile(walletPath)

	if err != nil {
		return nil, err
	}

	var wallet []string
	err = yaml.Unmarshal(content, &wallet)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
