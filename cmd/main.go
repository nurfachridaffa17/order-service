package main

import (
	"order-service/config"
	"order-service/internal/pkg/db"
	"order-service/internal/pkg/logging"
	"order-service/internal/pkg/migration"
	"os"

	"os/signal"
	"syscall"

	grpc_server "order-service/internal/api/grpc-server"
	"order-service/internal/api/http"
)

func main() {
	conf := config.NewEnv()
	conf.Load()

	logging.Init()

	db.Init()

	migration.Init()

	// Initialize GRPC
	go grpc_server.RunServer(db.DBManager())
	go http.Init(db.DBManager())

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logging.Log.Info("Shutdown Server ...")

}
