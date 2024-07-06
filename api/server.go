package main

import (
	"flag"
	"github.com/jongyunha/advance-go-web-application/api/apis"
	"github.com/jongyunha/advance-go-web-application/api/app"
	"github.com/jongyunha/advance-go-web-application/api/core"
	"go.uber.org/zap"
	"log"
)

func main() {
	var stage string
	flag.StringVar(&stage, "stage", "stage of the application", "development")

	application, err := app.New(core.Stage(stage))
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}

	err = apis.Serve(application)
	if err != nil {
		core.Logger.Error("failed to serve", zap.Error(err))
	}
}
