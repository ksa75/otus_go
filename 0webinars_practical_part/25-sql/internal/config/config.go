package config

import "github.com/BurntSushi/toml"

func Read(fpath string) (c Config, err error) {
	_, err = toml.DecodeFile(fpath, &c)
	return
}

type Config struct {
	PSQL PSQLConfig
}

type PSQLConfig struct {
	DSN string
}
