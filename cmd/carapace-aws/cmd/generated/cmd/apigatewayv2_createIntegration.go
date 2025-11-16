package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createIntegrationCmd = &cobra.Command{
	Use:   "create-integration",
	Short: "Creates an Integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createIntegrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_createIntegrationCmd).Standalone()

		apigatewayv2_createIntegrationCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_createIntegrationCmd.Flags().String("connection-id", "", "The ID of the VPC link for a private integration.")
		apigatewayv2_createIntegrationCmd.Flags().String("connection-type", "", "The type of the network connection to the integration endpoint.")
		apigatewayv2_createIntegrationCmd.Flags().String("content-handling-strategy", "", "Supported only for WebSocket APIs.")
		apigatewayv2_createIntegrationCmd.Flags().String("credentials-arn", "", "Specifies the credentials required for the integration, if any.")
		apigatewayv2_createIntegrationCmd.Flags().String("description", "", "The description of the integration.")
		apigatewayv2_createIntegrationCmd.Flags().String("integration-method", "", "Specifies the integration's HTTP method type.")
		apigatewayv2_createIntegrationCmd.Flags().String("integration-subtype", "", "Supported only for HTTP API AWS\\_PROXY integrations.")
		apigatewayv2_createIntegrationCmd.Flags().String("integration-type", "", "The integration type of an integration.")
		apigatewayv2_createIntegrationCmd.Flags().String("integration-uri", "", "For a Lambda integration, specify the URI of a Lambda function.")
		apigatewayv2_createIntegrationCmd.Flags().String("passthrough-behavior", "", "Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the requestTemplates property on the Integration resource.")
		apigatewayv2_createIntegrationCmd.Flags().String("payload-format-version", "", "Specifies the format of the payload sent to an integration.")
		apigatewayv2_createIntegrationCmd.Flags().String("request-parameters", "", "For WebSocket APIs, a key-value map specifying request parameters that are passed from the method request to the backend.")
		apigatewayv2_createIntegrationCmd.Flags().String("request-templates", "", "Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.")
		apigatewayv2_createIntegrationCmd.Flags().String("response-parameters", "", "Supported only for HTTP APIs.")
		apigatewayv2_createIntegrationCmd.Flags().String("template-selection-expression", "", "The template selection expression for the integration.")
		apigatewayv2_createIntegrationCmd.Flags().String("timeout-in-millis", "", "Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.")
		apigatewayv2_createIntegrationCmd.Flags().String("tls-config", "", "The TLS configuration for a private integration.")
		apigatewayv2_createIntegrationCmd.MarkFlagRequired("api-id")
		apigatewayv2_createIntegrationCmd.MarkFlagRequired("integration-type")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_createIntegrationCmd)
}
