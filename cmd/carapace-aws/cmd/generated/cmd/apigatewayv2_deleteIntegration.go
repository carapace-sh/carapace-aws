package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteIntegrationCmd = &cobra.Command{
	Use:   "delete-integration",
	Short: "Deletes an Integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_deleteIntegrationCmd).Standalone()

		apigatewayv2_deleteIntegrationCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_deleteIntegrationCmd.Flags().String("integration-id", "", "The integration ID.")
		apigatewayv2_deleteIntegrationCmd.MarkFlagRequired("api-id")
		apigatewayv2_deleteIntegrationCmd.MarkFlagRequired("integration-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteIntegrationCmd)
}
