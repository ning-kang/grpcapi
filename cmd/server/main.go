package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/ning-kang/grpcapi/internal"
	"github.com/ning-kang/grpcapi/protogen/golang/bookstore"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	hostString := os.Getenv("HOST")
	if hostString == "" {
		log.Fatal("HOST is not found in the environment file")
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment file")
	}

	host := hostString + ":" + portString

	// create new grpc server
	s := grpc.NewServer()

	// create new echo service
	bs := internal.NewBookStore()

	// register the new echo service to grpc server
	bookstore.RegisterBookStoreServer(s, bs)

	// listen to requests
	tl, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(fmt.Errorf("error starting tcp listener: %w", err))
	}

	// start listening
	s.Serve(tl)
}
