package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewaymanagementapi_postToConnectionCmd = &cobra.Command{
	Use:   "post-to-connection",
	Short: "Sends the provided data to the specified connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewaymanagementapi_postToConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewaymanagementapi_postToConnectionCmd).Standalone()

		apigatewaymanagementapi_postToConnectionCmd.Flags().String("connection-id", "", "The identifier of the connection that a specific client is using.")
		apigatewaymanagementapi_postToConnectionCmd.Flags().String("data", "", "The data to be sent to the client specified by its connection id.")
		apigatewaymanagementapi_postToConnectionCmd.MarkFlagRequired("connection-id")
		apigatewaymanagementapi_postToConnectionCmd.MarkFlagRequired("data")
	})
	apigatewaymanagementapiCmd.AddCommand(apigatewaymanagementapi_postToConnectionCmd)
}
