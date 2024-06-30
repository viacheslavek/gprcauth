package main

import (
	"github.com/viacheslavek/grpcauth/auth/internal/app"
	"github.com/viacheslavek/grpcauth/auth/internal/config"
	"github.com/viacheslavek/grpcauth/auth/internal/lib/logger"
)

// TODO: В proto вместо password надо хранить и отдавать passwordHash - исправить

func main() {
	cfg := config.MustLoad()
	lg := logger.SetupLogger(cfg.Env)

	application := app.New(lg, cfg.GRPC.Port, cfg.DB, cfg.TokenTTL)

	go func() {
		application.GRPCServer.MustRun()
	}()

	application.GracefulStop()
}
