package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Log struct {
	Path     string `yaml:"path" env-required:"true"`
	Filename string `yaml:"filename" env-required:"true"`
}

type Psql struct {
	Host          string `yaml:"host" env-required:"true"`
	Port          string `yaml:"port" env-required:"true"`
	Database      string `yaml:"database" env-required:"true"`
	Username      string `yaml:"username" env-required:"true"`
	Password      string `yaml:"password" env-required:"true"`
	PgPoolMaxConn int    `yaml:"pg_pool_max_conn" env-required:"true"`
}

type Redis struct {
	Addr     string `yaml:"addr" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Grpc struct {
	Host string `yaml:"host" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
}

type Storage struct {
	Psql  Psql  `yaml:"psql"`
	Redis Redis `yaml:"redis"`
}

type Config struct {
	IsDebug  *bool   `yaml:"is_debug" env-required:"true"`
	GrpcRepo Grpc    `yaml:"grpc_repo" env-required:"true"`
	Storage  Storage `yaml:"storage" env-required:"true"`
	Log      Log     `yaml:"log" env-required:"true"`
	FilePath string  `yaml:"file_path" env-required:"true"`
}

var _instance *Config
var _once sync.Once

func GetConfig() *Config {
	_once.Do(func() {

		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pathConfig := pwd + "/config.yml"
		fmt.Println(pathConfig)
		fmt.Println("read application configuration: pwd: ", pwd)

		_instance = &Config{}
		if err := cleanenv.ReadConfig(pathConfig, _instance); err != nil {
			help, _ := cleanenv.GetDescription(_instance, nil)
			fmt.Println(help)
			fmt.Println(err)
			panic("config.yml not correct")
		}
	})
	return _instance
}
