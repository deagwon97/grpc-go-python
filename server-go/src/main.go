package main

import (
	"encoding/json"
	"grpc-go/countries"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	pb "grpc-go/helloworld/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Server is implementation proto interface
type Server struct{}

// Search function responsible to get the Country information
func (Server) Search(ctx context.Context, request *countries.CountryRequest) (*countries.CountryResponse, error) {
	resp, err := http.Get("https://restcountries.com/v2/name/" + request.Name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data []countries.CountryResponse
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return nil, err
	}
	return &data[0], nil
}

// helloWorld server

type helloWorldServer struct {
	pb.UnimplementedGreeterServer
}

func (s *helloWorldServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	grpcServer := grpc.NewServer()
	var server Server
	countries.RegisterCountryServer(grpcServer, server)

	pb.RegisterGreeterServer(grpcServer, &helloWorldServer{})

	listen, err := net.Listen("tcp", "grpc-server-go:8000")
	if err != nil {
		log.Fatalf("could not listen to grpc-server-go:8000 %v", err)
	}
	log.Println("Server starting...")
	log.Println(grpcServer.Serve(listen))
}
