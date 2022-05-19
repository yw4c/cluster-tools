package cmd

import (
	"cluster-tools/app"
	"cluster-tools/c"

	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   string(c.HTTP),
	Short: "start http server",
	Long:  "http \n --port 8080",
	Run: func(cmd *cobra.Command, args []string) {
		server := app.NewHttpApp(port)
		server.Run(app.RegisterProvidersFunc)
	},
}
