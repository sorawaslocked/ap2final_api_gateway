package main

import (
	"github.com/sorawaslocked/ap2final_api_gateway/internal/config"
	"log"
)

func main() {
	// TODO: init config
	cfg := config.MustLoad()

	log.Println(cfg)
}
