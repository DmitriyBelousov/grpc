package adder

import (
	"context"
	"fmt"
	"log"

	"github.com/DmitriyBelousov/grpc/pkg/api"
	"google.golang.org/grpc/metadata"
)

const KekKey = "kek"

type GRPCServer struct{}

func (s *GRPCServer)Add(ctx context.Context,req *api.AddRequest)( *api.AddResponse,error){
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok{
		log.Fatal("get md error")
	}

	k := md.Get(KekKey)
	fmt.Println(k)
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
