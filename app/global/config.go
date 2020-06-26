package global

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

const defaultConfigFile = "../../../config.yaml"

func LoadConfig() {
	v := viper.New()
	v.SetConfigFile(defaultConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v\n", err)
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config file changed: %v\n", e.Name)
		if err := v.Unmarshal(&Config); err != nil {
			log.Printf("Fail to unmarshals the config into a Struct: %v\n", err)
		}
	})
	if err := v.Unmarshal(&Config); err != nil {
		log.Printf("Fail to unmarshals the config into a Struct: %v\n", err)
	}
}
