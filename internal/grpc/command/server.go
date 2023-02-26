package command

import (
	api "githab.com/grpc/server/proto"
)

type Server struct {
	api.UnimplementedGreeterServer
}

// func NewServer(a *api.UnimplementedGreeterServer) *server{
// 	return &server{
// 		api.UnimplementedGreeterServer: a,
// 	}
// }

// func NewCommander(
// 	bot *tgbotapi.BotAPI,
// 	productService *product.Service,
// ) *Commander {
// 	return &Commander{
// 		bot:            bot,
// 		productService: productService,
// 	}
// }
