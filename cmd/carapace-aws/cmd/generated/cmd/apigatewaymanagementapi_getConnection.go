package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewaymanagementapi_getConnectionCmd = &cobra.Command{
	Use:   "get-connection",
	Short: "Get information about the connection with the provided id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewaymanagementapi_getConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewaymanagementapi_getConnectionCmd).Standalone()

		apigatewaymanagementapi_getConnectionCmd.Flags().String("connection-id", "", "")
		apigatewaymanagementapi_getConnectionCmd.MarkFlagRequired("connection-id")
	})
	apigatewaymanagementapiCmd.AddCommand(apigatewaymanagementapi_getConnectionCmd)
}
