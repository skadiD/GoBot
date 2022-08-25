package config

import (
	"github.com/fexli/logger"
	"gobot/cmd/globals"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

// Create 配置文件创建
func Create(config interface{}) {
	text, err := yaml.Marshal(&config)
	if err != nil {
		logger.RootLogger.Warning(logger.WithContent("配置文件创建失败", err))
		err = nil
	}
	f, err := os.OpenFile("Config.yaml", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logger.RootLogger.Error(logger.WithContent("配置文件创建失败", err))
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		f.WriteAt(text, n)
		defer f.Close()
	}
}

// Read 配置文件读取
func Read() *globals.Common {
	rand.Seed(time.Now().Unix())
	config := new(globals.Common)
	if !FileExist("Config.yaml") {
		logger.RootLogger.Warning(logger.WithContent("配置文件不存在，自动生成中"))
		Create(config)
		logger.RootLogger.System(logger.WithContent("配置文件创建完毕，请及时修改"))
	}
	content, err := ioutil.ReadFile("Config.yaml")
	if err != nil {
		logger.RootLogger.Warning(logger.WithContent("读取 Config.yaml 出错", err))
	}
	if yaml.Unmarshal(content, &config) != nil {
		logger.RootLogger.Warning(logger.WithContent("解析 Config.yaml 出错"))
	}
	return config
}

// FileExist 文件是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
