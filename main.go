package main

import (
	"log"

	"github.com/nesar-ahmed/goecommerce/config"
	"github.com/nesar-ahmed/goecommerce/internal/api"
)


func main()  {
	cfg, err := config.SetupEnv()
	
	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.StartServer(cfg)
}