package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_putIntegrationCmd = &cobra.Command{
	Use:   "put-integration",
	Short: "Sets up a method's integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_putIntegrationCmd).Standalone()

	apigateway_putIntegrationCmd.Flags().String("cache-key-parameters", "", "A list of request parameters whose values API Gateway caches.")
	apigateway_putIntegrationCmd.Flags().String("cache-namespace", "", "Specifies a group of related cached parameters.")
	apigateway_putIntegrationCmd.Flags().String("connection-id", "", "The ID of the VpcLink used for the integration.")
	apigateway_putIntegrationCmd.Flags().String("connection-type", "", "The type of the network connection to the integration endpoint.")
	apigateway_putIntegrationCmd.Flags().String("content-handling", "", "Specifies how to handle request payload content type conversions.")
	apigateway_putIntegrationCmd.Flags().String("credentials", "", "Specifies whether credentials are required for a put integration.")
	apigateway_putIntegrationCmd.Flags().String("http-method", "", "Specifies the HTTP method for the integration.")
	apigateway_putIntegrationCmd.Flags().String("integration-http-method", "", "The HTTP method for the integration.")
	apigateway_putIntegrationCmd.Flags().String("passthrough-behavior", "", "Specifies the pass-through behavior for incoming requests based on the Content-Type header in the request, and the available mapping templates specified as the `requestTemplates` property on the Integration resource.")
	apigateway_putIntegrationCmd.Flags().String("request-parameters", "", "A key-value map specifying request parameters that are passed from the method request to the back end.")
	apigateway_putIntegrationCmd.Flags().String("request-templates", "", "Represents a map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.")
	apigateway_putIntegrationCmd.Flags().String("resource-id", "", "Specifies a put integration request's resource ID.")
	apigateway_putIntegrationCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_putIntegrationCmd.Flags().String("timeout-in-millis", "", "Custom timeout between 50 and 29,000 milliseconds.")
	apigateway_putIntegrationCmd.Flags().String("tls-config", "", "")
	apigateway_putIntegrationCmd.Flags().String("type", "", "Specifies a put integration input's type.")
	apigateway_putIntegrationCmd.Flags().String("uri", "", "Specifies Uniform Resource Identifier (URI) of the integration endpoint.")
	apigateway_putIntegrationCmd.MarkFlagRequired("http-method")
	apigateway_putIntegrationCmd.MarkFlagRequired("resource-id")
	apigateway_putIntegrationCmd.MarkFlagRequired("rest-api-id")
	apigateway_putIntegrationCmd.MarkFlagRequired("type")
	apigatewayCmd.AddCommand(apigateway_putIntegrationCmd)
}
