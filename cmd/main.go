package main

import (
	"github.com/JamshedJ/InstallmentPay/internal/handler"
	"github.com/JamshedJ/InstallmentPay/pkg/glog"
)

func main() {
	logger := glog.NewLogger()

	err := handler.Run(":8080")
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to start the server")
	}
}