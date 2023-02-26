package command

import (
	"context"
	"log"

	"githab.com/grpc/server/migration/database"
	api "githab.com/grpc/server/proto"
)

// type User struct {
// 	id         string
// 	age        string
// 	first_name string
// 	last_name  string
// 	email      string
// }

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &api.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *Server) GetUsers(ctx context.Context, in *api.Empty) (*api.UsersResult, error) {
	//con := database.NewConnection()
	// con2:= database.
	// database.NewConnection()
	// user := con.get_all_users()

	// data := database.NewService()
	// user := database.NewConnection()
	// // user.

	c := database.NewConnection()
	users := c.Get_all_users()

	// u := database.User{
	// 	id: "1", age: "2", first_name: "Artur", last_name: "Mish", email: "Artur@mail.ru",
	// }

	// num1 := User{123421

	//new_users := [User{	},{},{}]
	//line := []string{"str1", "str3", "str4", "str5", "str5"}
	//line2 := "str"
	return &api.UsersResult{Data: users}, nil
}

func (s *Server) CreateUser(ctx context.Context, in *api.User) (*api.Empty, error) {
	c := database.NewConnection()
	c.Create_New_User(in.Id, in.Age, in.FirstName, in.LastName, in.Email)
	return &api.Empty{}, nil
}

func (s *Server) DeletUser(ctx context.Context, in *api.Email) (*api.Empty, error) {
	c := database.NewConnection()
	c.Delete_User(in.Message)

	return &api.Empty{}, nil
}
