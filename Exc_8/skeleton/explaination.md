# What did we do?

We set up a gRPC server.

Meaning an alternative to REST where we don't use http and JSON (JSON sucks, so yay :3).

How? ->

## 1. Define .proto file

### 1.1 Messages

In the .proto file we define our data structs (how do the messages look, i.e. What makes up a "Drink"), the number assigned is the priority of in the bit string, i.e. just count up from 1

The structs are called message.

```proto
message Drink {
  int32 id = 1;
  string name = 2;
  float price = 3;
}
````

#### Lists

To make a list/slice, (i.e. []*Drink -> in go) we use keyword `repeat` ->

```proto
message Drinks {
  repeated Drink drinks = 1;
}
```

### 1.2 Services

In the services we define the functions the the gRPC server calls.
i.e. the names of the functions we will call on the server object in the go code both when sending a message and returning the code.

So each function is implimented on the server (serves the requests) and the client recives the answer of the server.
e.g.

```proto
service OrderService {
    rpc GetDrinks(google.protobuf.Empty) returns (Drinks) {}
    ...
}
```

- rcv -> keyword for defining function
- GetDrinks -> function name callable in go
- (google.protobuf.Empty) -> the message passed to the function; here nothing is passed
- return -> what the function returns a message or bool,...
- (Drinks) -> the function returns a [Drinks](#lists) message struct

## 2. compile

Using the script or the command

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/orders.proto
```

The Protoc compiler automatically compiles the requred code for sending/recieving into go files with the command in the pb folder.
These files do not need to be edited.

## 3. Making the server

In the [server/grpc_server.go](server/grpc_server.go) the following struct is provided:

```go
type GRPCService struct {
    pb.UnimplementedOrderServiceServer;
}
```

It represents our server and is the struct for which we have to impliment the functions we defined in the .proto file, which where compiled into the `pb` folder and are already imported.

### 3.1 Saving state

When we want to save something inside the server in memory we also do it by adding it to the server i.e.

```go
type GRPCService struct {
    pb.UnimplementedOrderServiceServer;
    drinks * pb.Drinks;
    orders * pb.Orders;
}
```

The definitions for the datatypes are from the generated file as can be seen by the `pb.` prefix before `pb.Drinks`, the Drinks message in .proto was translated by the proto compiler into a go struct.

### 3.2 Functions

The functions are implimented for the struct in the same file, they are the logic the server will execute when asked for said message through the client.

The functions can be found and copied from the generated file [pb/orders_grpc.pb.go](pb/orders_grpc.pb.go#49) at roughly line 49 but the `orderServiceClient` or equivalent will have to be changed to the struct of the server.

example:

```go
func (s * GRPCService) GetOrders(context.Context, *emptypb.Empty) (*pb.Orders, error){
    if s.orders == nil {
        return nil, errors.New("no orders")
    }
    return s.orders, nil
}
```

parts:

- func -> marks it as function
- (s * GRPCService) -> marks it as a method of the GRPCService struct
  - the * means we take the struct as a pointer and directly modify data inside of the object itself and not on a copy
- GetOrders -> Name of the function, from the .proto rpc function defintions
- (context.Context, *emptypb.Empty) -> the parameters of the function
  - the context is automatically added
  - the other parts consists of the messages we decided to pass it in the .proto defition
- (*pb.Orders, error) -> return values sent to the client on complition of the function

## 4 Making the client

The RUN method is implimented on the pointer of `GrpcClient` which we can use to call our functions and work with their return values.

example:

```go
func (c *GrpcClient) Run() error {
    drinks, err := c.client.GetDrinks(context.Background(), &emptypb.Empty{});
    for _, drink := range drinks.Drinks...
    ...
}
```

Calling functions here will query the server which is started when executing the `main.go` with `go run ./main.go`.

## 5 starting the server

In the `main.go` we only need to import the server and client files we finished in steps 3 and 4 and create an object of the defined server and clienit structs.

```go
import (
    ...
    "exc8/server"
)
func main() {
    ...
    err := server.StartGrpcServer();
    ...
}
```

Now when executing the `main.go` it will start our server, wait than start the client which when executing the `RUN()` function will query our server and thus our gRPC client and server setup are complete.

Yay :)
