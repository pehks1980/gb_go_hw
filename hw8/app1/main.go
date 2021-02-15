package main

import (
	"fmt"
	"github.com/pehks1980/gb_go_hw/hw8/app1/yamlcfg"
	"log"
)

func main() {
	// declare config struct
	var c yamlcfg.Conf
	// load config
	c.GetConf("conf.yaml")
	// validate config
	confStr := c.ValidateConfig()

	if !c.Valid {
		log.Fatalf("validate config error: %s", confStr)
	}

	fmt.Printf("--- yaml config:\n%v\n\n", c)


}
