package main

import (
	"fmt"
	"log"

	"video-feed/internal/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatalf("init app failed: %v", err)
	}

	addr := fmt.Sprintf(":%d", application.Config.App.Port)
	if err := application.Engine.Run(addr); err != nil {
		log.Fatalf("run server failed: %v", err)
	}
}
