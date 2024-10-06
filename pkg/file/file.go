package file

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadFromYAML(path string, target any) error {
	yf, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yf, target)
}
