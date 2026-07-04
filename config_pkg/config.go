package configpkg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int32  `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

var Host, Username, Password string
var Port int32

func Configpkg() {
	f, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println("Error while reading config.yaml", err)
		os.Exit(1)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Error while decoding config")
		os.Exit(1)
	}

	Host = cfg.Server.Host
	Port = cfg.Server.Port
	Username = cfg.Database.Username
	Password = cfg.Database.Password
}
