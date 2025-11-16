package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getIntegrationResponseCmd = &cobra.Command{
	Use:   "get-integration-response",
	Short: "Gets an IntegrationResponses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getIntegrationResponseCmd).Standalone()

	apigatewayv2_getIntegrationResponseCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getIntegrationResponseCmd.Flags().String("integration-id", "", "The integration ID.")
	apigatewayv2_getIntegrationResponseCmd.Flags().String("integration-response-id", "", "The integration response ID.")
	apigatewayv2_getIntegrationResponseCmd.MarkFlagRequired("api-id")
	apigatewayv2_getIntegrationResponseCmd.MarkFlagRequired("integration-id")
	apigatewayv2_getIntegrationResponseCmd.MarkFlagRequired("integration-response-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getIntegrationResponseCmd)
}
