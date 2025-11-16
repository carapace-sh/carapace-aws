package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteIntegrationResponseCmd = &cobra.Command{
	Use:   "delete-integration-response",
	Short: "Deletes an IntegrationResponses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteIntegrationResponseCmd).Standalone()

	apigatewayv2_deleteIntegrationResponseCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_deleteIntegrationResponseCmd.Flags().String("integration-id", "", "The integration ID.")
	apigatewayv2_deleteIntegrationResponseCmd.Flags().String("integration-response-id", "", "The integration response ID.")
	apigatewayv2_deleteIntegrationResponseCmd.MarkFlagRequired("api-id")
	apigatewayv2_deleteIntegrationResponseCmd.MarkFlagRequired("integration-id")
	apigatewayv2_deleteIntegrationResponseCmd.MarkFlagRequired("integration-response-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteIntegrationResponseCmd)
}
