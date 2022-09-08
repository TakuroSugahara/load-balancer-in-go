package main

import (
	"fmt"
	"log"

	"github.com/TakuroSugahara/load-balancer/config"
	loadbalancer "github.com/TakuroSugahara/load-balancer/lb"
	"github.com/TakuroSugahara/load-balancer/rp"
)

func main() {
	cfg := config.NewConfig()
	fmt.Printf("hello %+v", cfg)

	lb := loadbalancer.New(cfg.Backends)
	rp := rp.New(cfg.Proxy.Port, lb.Handler)
	if err := rp.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
