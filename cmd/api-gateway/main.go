package main

import (
	"github.com/sorawaslocked/ap2final_api_gateway/internal/app"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/config"
	"github.com/sorawaslocked/ap2final_base/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	application, err := app.New(cfg, log)
	if err != nil {
		return
	}

	application.Run()
}
