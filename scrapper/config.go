package scrapper

import (
	"errors"
	"fmt"
	"os"
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
