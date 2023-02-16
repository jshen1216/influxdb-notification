package utils

import (
	"fmt"
	"influx-alert-api/global"
	"log"
	"strings"

	"github.com/spf13/viper"
)

func LoadEnvironment() {
	loadConfigFile()
	viperConfigToModel()
}

func loadConfigFile() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("沒有發現 config.yml，改抓取環境變數")
			viper.AutomaticEnv()
			viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		} else {
			// 有找到 config.yml 但是發生了其他未知的錯誤
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func viperConfigToModel() {

	if err := viper.Unmarshal(&global.EnvConfig); err != nil {
		log.Fatal("Error to parse env parameter")
	}
}
