package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ning-kang/grpcapi/internal"
	"github.com/ning-kang/grpcapi/protogen/golang/bookstore"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	// load application settings
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(fmt.Errorf("error: config File is not found: %w", err))
		} else {
			panic(fmt.Errorf("error: failed reading config file: %w", err))
		}
	}
	host := viper.GetString("server.host") + ":" + viper.GetString("server.port")

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
