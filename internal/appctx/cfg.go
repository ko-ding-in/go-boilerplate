package appctx

import (
	"fmt"
	"github.com/ko-ding-in/go-boilerplate/pkg/file"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
	"time"
)

const (
	configPath = "config/"
)

var (
	cfgOnce sync.Once
	_cfg    *Config
)

type (
	Config struct {
		App    App    `yaml:"app" json:"app"`
		Logger Logger `yaml:"log" json:"log"`
	}

	App struct {
		Name         string        `yaml:"name" json:"name"`
		Port         int           `yaml:"port" json:"port"`
		Debug        bool          `yaml:"debug" json:"debug"`
		Timezone     string        `yaml:"timezone" json:"timezone"`
		Env          string        `yaml:"env" json:"env"`
		ReadTimeout  time.Duration `yaml:"read_timeout" json:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout" json:"write_timeout"`
	}

	Logger struct {
		Level string `yaml:"level" json:"level"`
	}
)

func NewConfig() *Config {
	cfgPath := []string{configPath}
	cfgOnce.Do(func() {
		c, err := readConfig("app.yaml", cfgPath...)
		if err != nil {
			log.Fatal("failed to load config")
		}
		_cfg = c
	})
	return _cfg
}

func readConfig(configFile string, configPaths ...string) (*Config, error) {
	var (
		cfg  *Config
		errs []error
	)

	for _, path := range configPaths {
		cfgPath := fmt.Sprint(path, configFile)
		if err := file.ReadFromYAML(cfgPath, &cfg, func(s string) ([]byte, error) {
			return os.ReadFile(path)
		}, func(bytes []byte, a any) error {
			return yaml.Unmarshal(bytes, a)
		}); err != nil {
			errs = append(errs, fmt.Errorf("file %s error %s", cfgPath, err.Error()))
			continue
		}
		break
	}

	if cfg == nil {
		return nil, fmt.Errorf("file config parse error %v", errs)
	}

	return cfg, nil
}
