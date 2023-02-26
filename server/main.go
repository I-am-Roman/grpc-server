package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"githab.com/grpc/server/internal/grpc/command"
	api "githab.com/grpc/server/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// // SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &api.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

// func (s *server) GetUsers(ctx context.Context, in *api.Empty) (*api.UsersResult, error) {
// 	users := []string{"Ivan", "Petr", "Sigor", "Maluta", "Grey"}
// 	return &api.UsersResult{Users: users}, nil
// }

// type server struct {
// 	api.UnimplementedGreeterServer
// }

// type server struct {
// 	api.UnimplementedGreeterServer
// }

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//some_var :=  command.server{}

	api.RegisterGreeterServer(s, &command.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
