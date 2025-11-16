package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewaymanagementapi_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Delete the connection with the provided id.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewaymanagementapi_deleteConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewaymanagementapi_deleteConnectionCmd).Standalone()

		apigatewaymanagementapi_deleteConnectionCmd.Flags().String("connection-id", "", "")
		apigatewaymanagementapi_deleteConnectionCmd.MarkFlagRequired("connection-id")
	})
	apigatewaymanagementapiCmd.AddCommand(apigatewaymanagementapi_deleteConnectionCmd)
}
