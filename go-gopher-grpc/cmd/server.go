package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/1055373165/learn-go-by-example/go-gopher-grpc/pkg/gopher"
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

// server is used to implement gopher.GopherServer.
type Server struct {
	pb.UnimplementedGopherServer
}

type Gopher struct {
	URL string `json:"url"`
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the Schema gRPC server",

	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()

		// Register services
		pb.RegisterGopherServer(grpcServer, &Server{})

		log.Printf("GRPC server listening on %v", lis.Addr())

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

// GetGopher implements gopher.GopherServer
func (s *Server) GetGopher(ctx context.Context, req *pb.GopherRequest) (*pb.GopherReply, error) {
	res := &pb.GopherReply{}

	// Check request
	if req == nil {
		fmt.Println("request must not be nil")
		return res, xerrors.Errorf("request must not be nil")
	}

	if req.Name == "" {
		fmt.Println("name must not be empty in the request")
		return res, xerrors.Errorf("name must not be empty in the request")
	}

	log.Printf("Received: %v", req.GetName())

	// server process
	response := req.Name + ": RPC Call Success" + " return reply test"
	res.Message = response
	return res, nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
