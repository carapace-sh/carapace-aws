package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewaymanagementapiCmd = &cobra.Command{
	Use:   "apigatewaymanagementapi",
	Short: "The Amazon API Gateway Management API allows you to directly manage runtime aspects of your deployed APIs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewaymanagementapiCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewaymanagementapiCmd).Standalone()

	})
	rootCmd.AddCommand(apigatewaymanagementapiCmd)
}
