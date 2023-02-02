package engine

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func PrepareGrpcServer() (*grpc.Server, net.Listener) {
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")

	lis, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return grpc.NewServer(), lis
}

func RunGrpcServer(s *grpc.Server, lis net.Listener) {
	log.Printf("cmd listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
