package config

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger(stage Stage) error {
	if stage == Production {
		logger, err := zap.NewProduction()
		if err != nil {
			return err
		}
		Logger = logger
	} else {
		logger, err := zap.NewDevelopment()
		if err != nil {
			return err
		}

		Logger = logger
	}

	return nil
}
