package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/go-yaml/yaml"
)

// html DOM 查询配置
type Query struct {
	Title         []string `yaml:"title"`
	Content       []string `yaml:"content"`
	Next          []string `yaml:"next"`
	NextWithIndex bool     `yaml:"next_with_index"`
	NextIndex     uint     `yaml:"next_index"`
}

// 配置文件结构
type Conf struct {
	Base       string `yaml:"base"`
	Purl       string `yaml:"purl"`
	AppendMode bool   `yaml:"append_mode"`
	Encode     string `yaml:"encode"`
	Query      `yaml:"query"`
}

var configFile string
var useConf bool = false

// 判断文件是否存在
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

// 使用指定配置文件
func (conf *Conf) UseFile(cf string) {
	if IsFileExist(cf) {
		configFile = cf
	} else {
		fmt.Println("配置文件不存在")
		return
	}
	conf.parse(configFile)
}

// 使用默认配置文件
func (conf *Conf) GetConf() {
	configFile = defaultConfigFile()
	if !IsFileExist(configFile) {
		fmt.Println("配置文件不存在")
		return
	}
	conf.parse(configFile)
}

// 查找默认的配置文件
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

// 解析配置文件
func (conf *Conf) parse(configFile string) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("ioutil.ReadFile err:%s\n", err)
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
}
