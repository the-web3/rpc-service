package services

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"

	"github.com/the-web3/rpc-service/protobuf/wallet"
)

const MaxRecvMessageSize = 1024 * 1024 * 300

type RpcServerConfig struct {
	GrpcHostname string
	GrpcPort     int
}

type RpcServer struct {
	*RpcServerConfig
	wallet.UnimplementedWalletServiceServer
}

func NewRpcServer(config *RpcServerConfig) (*RpcServer, error) {
	return &RpcServer{
		RpcServerConfig: config,
	}, nil
}

func (s *RpcServer) Start() error {
	go func(s *RpcServer) {
		addr := fmt.Sprintf("%s:%d", s.GrpcHostname, s.GrpcPort)
		fmt.Println("start rpc server", "addr", addr)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			fmt.Println("Could not start tcp listener. ")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(
				nil,
			),
		)
		reflection.Register(gs)

		wallet.RegisterWalletServiceServer(gs, s)

		fmt.Println("Grpc info", "port", s.GrpcPort, "address", listener.Addr())
		if err := gs.Serve(listener); err != nil {
			fmt.Println("Could not GRPC server")
		}
	}(s)
	return nil
}
