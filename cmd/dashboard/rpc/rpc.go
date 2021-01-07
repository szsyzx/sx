package rpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/XOS/Probe/proto"
	rpcService "github.com/XOS/Probe/service/rpc"
)

// ServeRPC ...
func ServeRPC(port uint) {
	server := grpc.NewServer()
	pb.RegisterProbeServiceServer(server, &rpcService.ProbeHandler{
		Auth: &rpcService.AuthHandler{},
	})
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
