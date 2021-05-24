package server

import (
	"context"
	"fmt"
	"github.com/xuedi/starraid/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HandlerServer struct {
	api.UnimplementedHandlerServer
}

func (n *HandlerServer) Login(ctx context.Context, login *api.LoginRequest) (*api.LoginResponse, error) {
	log.Printf("Received Login (%s/%s)", login.GetUsername(), login.GetPassword())
	return &api.LoginResponse{}, nil
}

func (n *HandlerServer) Action(ctx context.Context, action *api.ActionRequest) (*api.ActionResponse, error) {
	log.Printf("Received Action with token: %s", action.GetToken())
	return &api.ActionResponse{}, nil
}

func (n *HandlerServer) FetchInfos(ctx context.Context, object *api.ObjectRequest) (*api.ObjectResponse, error) {
	log.Printf("Received object request for: %s", object.GetObjectId())
	return &api.ObjectResponse{}, nil
}



type Network struct {
	objects Objects
	handler  HandlerServer
}

func (n *Network) RegisterObjects(objects Objects) {
	n.objects = objects
}

func (n *Network) load(address string, port int) {
	n.handler = HandlerServer{}

	fmt.Println("Listening on %s:%d", address, port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("START")
	grpcServer := grpc.NewServer()  // attach the Ping service to the server

	api.RegisterHandlerServer(grpcServer, &n.handler)

	fmt.Println("SERVE")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	fmt.Println("DONE")
}



//func (n *Network) HandleLoginRequest(api.LoginRequest) (*PingMessage, error) {
//	log.Printf("Receive message %s", in.Greeting)
//	return &PingMessage{Greeting: "bar"}, nil
//}