package config

import "gopkg.in/yaml.v2"

type Conf struct {
	Port       int    `yaml:"port"`
	APIURL     string `yaml:"api_url"`
	DBHost     string `yaml:"db_host"`
	DBPort     int    `yaml:"db_port"`
	DBName     string `yaml:"db_name"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
}

func Parse(bytes []byte) (*Conf, error) {
	c := new(Conf)
	return c, yaml.Unmarshal(bytes, &c)
}
