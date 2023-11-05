package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Gvegas12/social-network-ws-api/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)

	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.NewServer(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
