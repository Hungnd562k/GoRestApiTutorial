package configpkg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Version string `yaml:"version"`
	} `yaml:"app"`
	Server struct {
		Host string `yaml:"host"`
		Port int32  `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Address           string `yaml:"address"`
		Username          string `yaml:"user"`
		Password          string `yaml:"pass"`
		ConnectionTimeout string `yaml:"connection_timeout"`
	} `yaml:"database"`
}

var Version, Host, Address, Username, Password, ConnectionTimeout string
var Port int32

func Configpkg() {
	f, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println(err)
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

	Version = cfg.App.Version

	Host = cfg.Server.Host
	Port = cfg.Server.Port

	Address = cfg.Database.Address
	Username = cfg.Database.Username
	Password = cfg.Database.Password
	ConnectionTimeout = cfg.Database.ConnectionTimeout
}
