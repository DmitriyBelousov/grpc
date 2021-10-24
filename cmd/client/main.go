package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	adderSvc "github.com/DmitriyBelousov/grpc/pkg/adder"
	"github.com/DmitriyBelousov/grpc/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main(){
	flag.Parse()

	if flag.NArg() <2{
		log.Fatalln("min 2 args")
	}

	c, err :=grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	adder := api.NewAdderClient(c)

	md:= metadata.Pairs(adderSvc.KekKey, "777")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err :=adder.Add(ctx, &api.AddRequest{
		X: int32(x),
		Y: int32(y),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.GetResult())
}
