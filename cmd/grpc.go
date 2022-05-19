package cmd

import (
	"cluster-tools/app"
	"cluster-tools/c"
	"log"

	"github.com/spf13/cobra"
)

var grpcCmd = &cobra.Command{
	Use:   string(c.GRPC),
	Short: "start grpc server",
	Long:  "http \n --port 8080",
	Run: func(cmd *cobra.Command, args []string) {
		defer log.Println("grpc server is ready")
		server := app.NewGrpcApp(port)
		server.Run(app.RegisterProvidersFunc)
	},
}
