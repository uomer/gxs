package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/go-yaml/yaml"
)

type Conf struct {
	Base       string `yaml:"base"`
	Purl       string `yaml:"purl"`
	AppendMode bool   `yaml:"append_mode"`
	Query      struct {
		Title         []string `yaml:"title"`
		Content       []string `yaml:"content"`
		Next          []string `yaml:"next"`
		NextWithIndex bool     `yaml:"next_with_index"`
		NextIndex     uint     `yaml:"next_index"`
	} `yaml:"query"`
}

var Config = new(Conf)
var configFile string
var useConf bool = false

func IsFileExist(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return false
	}
	return true
}

func UseFile(cf string) *Conf {
	if IsFileExist(cf) {
		configFile = cf
	} else {
		configFile = defaultConfigFile()
	}
	parse(configFile)
	return Config

}

func GetConf() *Conf {
	configFile = defaultConfigFile()
	parse(configFile)
	return Config
}

func defaultConfigFile() string {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("currentUser err: %v\n", err)
	}
	homeDir := currentUser.HomeDir

	configPaths := []string{
		"./config.yaml",
		"/etc/gxs/config.yaml",
		homeDir + "/.gxs.yaml",
	}
	for _, configPath := range configPaths {
		if IsFileExist(configPath) {
			configFile = configPath
			break
		}
	}
	return configFile
}

func parse(configFile string) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("ioutil.ReadFile err:%s\n", err)
	}
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
}
