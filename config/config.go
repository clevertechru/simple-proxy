package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"

	"gopkg.in/yaml.v2"
)

type config struct {
	ServiceName string
	ServicePort int
	Upstream    []string
}

func (c *config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var aux struct {
		ServiceName string   `yaml:"service_name"`
		ServicePort string   `yaml:"service_port"`
		Upstream    []string `yaml:"upstream"`
	}

	if err := unmarshal(&aux); err != nil {
		return err
	}
	if aux.ServiceName == "" {
		return errors.New("config: invalid `service_name`")
	}
	if len(aux.Upstream) == 0 {
		return errors.New("config: invalid `upstream`")
	}
	// Test Kitchen stores the port as an string
	port, err := strconv.Atoi(aux.ServicePort)
	if err != nil {
		return errors.New("config: invalid `service_port`")
	}

	c.ServiceName = aux.ServiceName
	c.Upstream = aux.Upstream
	c.ServicePort = port
	return nil
}

func ReadConfig() (*config, error) {
	configFile := path.Join("config.yml")
	fmt.Printf("config file = %s\n", configFile)

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	var serviceConfig config
	if err := yaml.Unmarshal(data, &serviceConfig); err != nil {
		return nil, err
	}

	return &serviceConfig, nil
}
