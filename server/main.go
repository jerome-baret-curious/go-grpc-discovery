package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	jba "jerome.baret/animal"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	jba.UnimplementedAnimalDialogServer
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	jba.RegisterAnimalDialogServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) UniAnimal(ctx context.Context, in *jba.AnimalRequest) (*jba.AnimalResponse, error) {
	return &jba.AnimalResponse{Eat: "Here's your " + in.GetFood()}, nil
}

func (s *server) ServStreamAnimal(in *jba.AnimalRequest, ss jba.AnimalDialog_ServStreamAnimalServer) error {
	for i := 0; i < 5; i++ {
		if err := ss.Send(&jba.AnimalResponse{Eat: "Take this " + in.GetFood()}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) CliStreamAnimal(cs jba.AnimalDialog_CliStreamAnimalServer) error {
	for {
		_, err := cs.Recv()
		if err == io.EOF {
			return cs.SendAndClose(&jba.AnimalResponse{Eat: "Yummy"})
		}
		if err != nil {
			return err
		}
		fmt.Println("Got food")
	}
}

func (s *server) BiAnimal(ss jba.AnimalDialog_BiAnimalServer) error {
	for {
		in, err := ss.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		for i := 0; i < 5; i++ {
			if err := ss.Send(&jba.AnimalResponse{Eat: "As requested: " + in.GetFood()}); err != nil {
				return err
			}
		}
	}
}
