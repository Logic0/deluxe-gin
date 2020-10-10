package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Config *config

func init() {
	Config = &config{}
	env := os.Getenv("ENVIRONMENT")
	var ymlContent []byte
	var err error
	switch env {
	case "pro": // 生产环境
		ymlContent, err = ioutil.ReadFile("./config/app-pro.yml")
	case "dev": // 测试环境
		ymlContent, err = ioutil.ReadFile("./config/app-dev.yml")
	default:
		ymlContent, err = ioutil.ReadFile("./config/app-dev.yml")
		// panic("Environment variable not set(dev/pro)!")
	}

	if err != nil {
		panic("config file app.yml open failed:" + err.Error())
	}

	err = yaml.Unmarshal(ymlContent, Config)
	if err != nil {
		panic("config file app.yml parse failed:" + err.Error())
	}

	Config.Environment = env
}
