package main

import (
	"log"
	"net"

	"github.com/DmitriyBelousov/grpc/pkg/adder"
	"github.com/DmitriyBelousov/grpc/pkg/api"
	"google.golang.org/grpc"
)

func main(){
	s  := grpc.NewServer()
	adderSvc := &adder.GRPCServer{}
	api.RegisterAdderServer(s, adderSvc)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = s.Serve(l); err != nil{
		log.Fatal(err)
	}
}

