package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type Config struct {
	vn          *viper.Viper
	configPath  string
	stage       Stage
	ProjectCode string `mapstructure:"project_code"`
	Maintenance bool   `mapstructure:"maintenance"`
}

func (c *Config) Init(stage, cfgPath string) error {
	c.stage = parseStage(stage)
	c.configPath = cfgPath

	cfgName := fmt.Sprintf("config.%s", c.stage.String())

	vn := viper.New()
	vn.AddConfigPath(c.configPath)
	vn.SetConfigName(cfgName)
	if err := vn.ReadInConfig(); err != nil {
		return err
	}
	c.vn = vn

	if err := c.binding(); err != nil {
		return err
	}

	vn.WatchConfig()
	vn.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config change:", e.Name)
		if err := c.binding(); err != nil {
			return
		}
	})

	return nil
}

func (c *Config) binding() error {
	return c.vn.Unmarshal(&c)
}

func parseStage(s string) Stage {
	switch s {
	case "dev", "d", "development":
		return StageDev
	case "sit":
		return StageSIT
	}
	return StageDev
}
