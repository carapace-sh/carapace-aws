package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_putGatewayResponseCmd = &cobra.Command{
	Use:   "put-gateway-response",
	Short: "Creates a customization of a GatewayResponse of a specified response type and status code on the given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_putGatewayResponseCmd).Standalone()

	apigateway_putGatewayResponseCmd.Flags().String("response-parameters", "", "Response parameters (paths, query strings and headers) of the GatewayResponse as a string-to-string map of key-value pairs.")
	apigateway_putGatewayResponseCmd.Flags().String("response-templates", "", "Response templates of the GatewayResponse as a string-to-string map of key-value pairs.")
	apigateway_putGatewayResponseCmd.Flags().String("response-type", "", "The response type of the associated GatewayResponse")
	apigateway_putGatewayResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_putGatewayResponseCmd.Flags().String("status-code", "", "The HTTP status code of the GatewayResponse.")
	apigateway_putGatewayResponseCmd.MarkFlagRequired("response-type")
	apigateway_putGatewayResponseCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_putGatewayResponseCmd)
}
