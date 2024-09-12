package main

import (
	"net"
	"order-service/config"
	"order-service/internal/handler"
	"order-service/internal/pkg/db"
	"order-service/internal/pkg/logging"
	"order-service/internal/pkg/migration"
	"order-service/internal/repository"
	"order-service/internal/service"
	order_service "order-service/proto/order-service/proto/order"
	"os"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func main() {
	conf := config.NewEnv()
	conf.Load()

	logging.Init()

	db.Init()

	migration.Init()

	err := runGRPCServer(db.DBManager())
	if err != nil {
		logging.Log.Fatalf("Failed to start gRPC server: %v", err)
		os.Exit(1)
	}

}

// runGRPCServer starts the gRPC server
func runGRPCServer(db *gorm.DB) error {
	// Initialize repositories
	orderRepo := repository.NewOrderRepository(db)
	orderLineRepo := repository.NewOrderLineRepository(db)

	// Initialize services
	orderService := service.NewOrderService(orderRepo, orderLineRepo)

	// Initialize gRPC handler
	orderHandler := handler.NewOrderHandler(orderService)

	// Setup gRPC server
	grpcServer := grpc.NewServer()

	// Register the OrderService server with the handler
	order_service.RegisterOrderServiceServer(grpcServer, orderHandler)

	// Setup listener on the gRPC port
	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		return err
	}

	logging.Log.Printf("gRPC server listening on %s", os.Getenv("GRPC_PORT"))

	// Start the gRPC server
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
