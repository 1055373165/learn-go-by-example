/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/1055373165/learn-go-by-example/go-gopher-grpc/pkg/gopher"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:9000"
	defaultName = "dr-who"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "query the grpc server",
	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		client := pb.NewGopherClient(conn)
		var name string
		// Contact the server and print out its response.
		if len(os.Args) > 2 {
			name = os.Args[2]
		} else {
			name = defaultName
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := client.GetGopher(ctx, &pb.GopherRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("URL: %s", r.GetMessage())
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
