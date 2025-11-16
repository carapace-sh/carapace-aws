package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createIntegrationResponseCmd = &cobra.Command{
	Use:   "create-integration-response",
	Short: "Creates an IntegrationResponses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createIntegrationResponseCmd).Standalone()

	apigatewayv2_createIntegrationResponseCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_createIntegrationResponseCmd.Flags().String("content-handling-strategy", "", "Specifies how to handle response payload content type conversions.")
	apigatewayv2_createIntegrationResponseCmd.Flags().String("integration-id", "", "The integration ID.")
	apigatewayv2_createIntegrationResponseCmd.Flags().String("integration-response-key", "", "The integration response key.")
	apigatewayv2_createIntegrationResponseCmd.Flags().String("response-parameters", "", "A key-value map specifying response parameters that are passed to the method response from the backend.")
	apigatewayv2_createIntegrationResponseCmd.Flags().String("response-templates", "", "The collection of response templates for the integration response as a string-to-string map of key-value pairs.")
	apigatewayv2_createIntegrationResponseCmd.Flags().String("template-selection-expression", "", "The template selection expression for the integration response.")
	apigatewayv2_createIntegrationResponseCmd.MarkFlagRequired("api-id")
	apigatewayv2_createIntegrationResponseCmd.MarkFlagRequired("integration-id")
	apigatewayv2_createIntegrationResponseCmd.MarkFlagRequired("integration-response-key")
	apigatewayv2Cmd.AddCommand(apigatewayv2_createIntegrationResponseCmd)
}
