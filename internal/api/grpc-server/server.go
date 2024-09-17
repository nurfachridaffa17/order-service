package grpc_server

import (
	"net"
	"order-service/internal/api/grpc-server/handler"
	"order-service/internal/pkg/logging"
	"order-service/internal/repository"
	"order-service/internal/service"
	"os"

	order_service "order-service/proto/order-service/proto/order"

	"google.golang.org/grpc"

	"gorm.io/gorm"
)

func RunServer(db *gorm.DB) error {
	return runGRPCServer(db)
}

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
