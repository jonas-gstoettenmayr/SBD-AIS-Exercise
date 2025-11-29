package server

import (
	"context"
	"exc8/pb"
	"errors"

	// "fmt"
	// "log/slog"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer;
	drinks * pb.Drinks;
	orders * pb.Orders;

}

func (s * GRPCService) prepopulateDrinks() () {
	if s.drinks == nil {
		s.drinks = &pb.Drinks{};
		d1 := pb.Drink {
			Id: 1,
			Name: "The bob",
			Price: 2.99,
		}
		d2 := pb.Drink {
			Id: 2,
			Name: "The drop",
			Price: 3.99,
		}
		d3 := pb.Drink {
			Id: 3,
			Name: "The tables",
			Price: 9.99,
		}
		s.drinks.Drinks = append(s.drinks.Drinks, &d1)
		s.drinks.Drinks = append(s.drinks.Drinks, &d2)
		s.drinks.Drinks = append(s.drinks.Drinks, &d3)
	}
}

func (s * GRPCService) GetDrinks(context.Context, *emptypb.Empty) (*pb.Drinks, error) {
	s.prepopulateDrinks()
	return s.drinks,  nil
}
func (s * GRPCService) GetOrders(context.Context, *emptypb.Empty) (*pb.Orders, error){
	if s.orders == nil {
		return nil, errors.New("no orders")
	}
	return s.orders, nil
}
func (s * GRPCService) OrderDrink(con context.Context, order *pb.Order) (*wrapperspb.BoolValue, error) {
	if s.orders == nil{
		s.orders = &pb.Orders{}
	}
	for _, old_order := range s.orders.Orders {
		if(old_order.Drink.Name == order.Drink.Name) {
			old_order.Amount += order.Amount
			return wrapperspb.Bool(true), nil
		}
	}
	s.orders.Orders = append(s.orders.Orders, order)
	return wrapperspb.Bool(true), nil
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}
