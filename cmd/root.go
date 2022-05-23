package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().IntVar(&port, "port", 8081, "server port")
}

var (
	port int
)

var rootCmd = &cobra.Command{
	SilenceUsage: true,
	Short:        "Start Server",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Println("start with command " + cmd.Use)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit
		log.Println("Shutting down server...")
	},
}

func Execute() {
	rootCmd.AddCommand(
		httpCmd,
		grpcCmd,
	)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
		return
	}
}
