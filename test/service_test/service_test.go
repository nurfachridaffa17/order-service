package test

import (
	"context"
	"log"
	"order-service/proto/order-service/proto/order"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestCreateOrderWithLines(t *testing.T) {
	conn, err := grpc.Dial("localhost:8111", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := order.NewOrderServiceClient(conn)

	req := &order.OrderRequest{
		TableId:   1,
		Total:     200.0,
		CreatedBy: 1,
		OrderLines: []*order.OrderLineRequest{
			{MenuId: 1, Quantity: 20, Price: 10.0},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateOrder(ctx, req)
	if err != nil {
		t.Fatalf("could not create order: %v", err)
	}

	assert.NotNil(t, res)
}
