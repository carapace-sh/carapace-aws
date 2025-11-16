package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateIntegrationResponseCmd = &cobra.Command{
	Use:   "update-integration-response",
	Short: "Updates an IntegrationResponses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateIntegrationResponseCmd).Standalone()

	apigatewayv2_updateIntegrationResponseCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_updateIntegrationResponseCmd.Flags().String("content-handling-strategy", "", "Supported only for WebSocket APIs.")
	apigatewayv2_updateIntegrationResponseCmd.Flags().String("integration-id", "", "The integration ID.")
	apigatewayv2_updateIntegrationResponseCmd.Flags().String("integration-response-id", "", "The integration response ID.")
	apigatewayv2_updateIntegrationResponseCmd.Flags().String("integration-response-key", "", "The integration response key.")
	apigatewayv2_updateIntegrationResponseCmd.Flags().String("response-parameters", "", "A key-value map specifying response parameters that are passed to the method response from the backend.")
	apigatewayv2_updateIntegrationResponseCmd.Flags().String("response-templates", "", "The collection of response templates for the integration response as a string-to-string map of key-value pairs.")
	apigatewayv2_updateIntegrationResponseCmd.Flags().String("template-selection-expression", "", "The template selection expression for the integration response.")
	apigatewayv2_updateIntegrationResponseCmd.MarkFlagRequired("api-id")
	apigatewayv2_updateIntegrationResponseCmd.MarkFlagRequired("integration-id")
	apigatewayv2_updateIntegrationResponseCmd.MarkFlagRequired("integration-response-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateIntegrationResponseCmd)
}
