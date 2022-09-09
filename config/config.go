package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/TakuroSugahara/load-balancer/backend"
)

type (
	Proxy struct {
		Port string `json:"port"`
	}
	Config struct {
		Proxy    Proxy            `json:"proxy"`
		Backends backend.Backends `json:"backends"`
	}
)

var cfg *Config

func NewConfig() *Config {
	if cfg != nil {
		return cfg
	}
	cfg = &Config{}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("cannot get wd", err.Error())
	}
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/config/config.json", wd))
	if err != nil {
		log.Fatal("cannot read config", err.Error())
	}
	if err := json.Unmarshal(data, cfg); err != nil {
		log.Fatal("cannot unmarshal config", err.Error())
	}
	return cfg
}
