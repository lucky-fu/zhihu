package config

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"
	"zhihu.com/m/define"
)

var (
	container = &sync.Map{}
)

func absPath(path string) string {
	base, _ := os.Getwd()
	return base + "/" + path + ".yml"
}

func read(relativePath string) ([]byte, error) {
	path := absPath(relativePath)

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	return ioutil.ReadFile(path)
}

func loadYaml(path string) (interface{}, error) {
	var config interface{}
	content, err := read(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	return config, nil
}

func get(path string) (interface{}, error) {
	var (
		value interface{}
		err   error
		ok    bool
	)

	value, ok = container.Load(path)

	if !ok {
		value, err = loadYaml(path)
		if err != nil {
			return nil, err
		}
	}

	container.Store(path, value)

	return value, err
}

// GetRedisConfig ...
func GetRedisConfig(path string) (*define.ConfigRedis, error) {
	var (
		retConfig define.ConfigRedis
	)

	rawConfig, err := get(path)
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(rawConfig, &retConfig)

	return &retConfig, err
}

// GetDBConfig ...
func GetDBConfig(path string) (*define.ConfigDB, error) {
	var (
		retConfig define.ConfigDB
	)

	rawConfig, err := get(path)
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(rawConfig, &retConfig)

	return &retConfig, err
}
