package config

import (
	"gobot/cmd/globals"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Create 配置文件创建
func Create(config interface{}) {
	text, err := yaml.Marshal(&config)
	if err != nil {
		globals.Logger().Warn("创建配置文件出错").Run()
		err = nil
	}
	f, err := os.OpenFile("Config.yaml", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		globals.Logger().Warn("创建配置文件出错").Run()
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		f.WriteAt(text, n)
		defer f.Close()
	}
}

// Read 配置文件读取
func Read() *globals.Common {
	config := new(globals.Common)
	if !FileExist("Config.yaml") {
		globals.Logger().Warn("配置文件不存在，自动生成中").Run()
		Create(config)
		globals.Logger().Info("配置文件创建完毕，请及时修改").Run()
	}
	content, err := ioutil.ReadFile("Config.yaml")
	if err != nil {
		globals.Logger().Warn("读取Config.yaml出错").Run()
	}

	if yaml.Unmarshal(content, &config) != nil {
		globals.Logger().Warn("解析Config.yaml出错")
	}
	return config
}

// FileExist 文件是否存在
func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
