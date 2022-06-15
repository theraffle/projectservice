package server

import (
	"fmt"
	"github.com/theraffle/projectservice/src/genproto/pb"
	"github.com/theraffle/projectservice/src/server/projectservice"
	"google.golang.org/grpc"
	"net"
	"os"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	log = logf.Log.WithName("server")
)

// Server is an interface of server
type Server interface {
	Start(string)
}

type projectServiceServer struct {
	grpcServer     *grpc.Server
	projectService *projectservice.ProjectService
}

// New returns new user service gRPC server
func New() Server {
	server := new(projectServiceServer)
	server.grpcServer = grpc.NewServer()
	server.projectService = &projectservice.ProjectService{}

	return server
}

// Start starts gRPC server on input port
func (s *projectServiceServer) Start(port string) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error(err, "cannot launch gRPC server")
		os.Exit(1)
	}

	pb.RegisterProjectServiceServer(s.grpcServer, s.projectService)
	if err := s.grpcServer.Serve(l); err != nil {
		log.Error(err, "cannot launch gRPC server")
		os.Exit(1)
	}
}
