package client

import (
	"context"
	"exc8/pb"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Run() error {
	// 1. List drinks
	if drinks, err := c.client.GetDrinks(context.Background(), &emptypb.Empty{}); err == nil{
		fmt.Printf("Requesting drinks 🍹🍺☕\n")
		fmt.Printf("Available drinks:\n")
		for _, drink := range drinks.Drinks{
			fmt.Printf("\t>id: %d - name: %s - price: %.2f\n", drink.Id, drink.Name, drink.Price)
		}

		// 2. Order a few drinks
		fmt.Printf("Ordering drinks 👨‍🍳⏱️🍻🍻\n")
		for i, drink := range drinks.Drinks{
			fmt.Printf("\t>ordering: %d x %s -> ", i+1, drink.Name)
			succeded, err := c.client.OrderDrink(context.Background(), &pb.Order{Id: int32(i), Amount: int32(i+1), Drink: drink})
			fmt.Printf("succed: %t, err: %s\n", succeded.Value, err)
		}

		// 3. Order more drinks
		fmt.Printf("Ordering another round of drinks 👨‍🍳⏱️🍻🍻\n")
		for i, drink := range drinks.Drinks{
			fmt.Printf("\t>ordering: %d x %s ->", 5, drink.Name)
			succeded, err := c.client.OrderDrink(context.Background(), &pb.Order{Id: int32(i+20), Amount: int32(5), Drink: drink})
			fmt.Printf("succed: %t, err: %s\n", succeded.Value, err)
		}
	}
	
	// 4. Get order total
	if orders, err := c.client.GetOrders(context.Background(), &emptypb.Empty{}); err == nil{
		fmt.Printf("Getting the bill 💹💹💹\n")
		var total float32 = 0
		for _, order := range orders.Orders {
			price := float32(order.Amount) * order.Drink.Price
			total += price; 
			fmt.Printf("\t>Total: %d x %s = %.2f\n", order.Amount, order.Drink.Name, price)
		}
		fmt.Printf("Total: %.2f\n", total)
	}

	return nil
}
