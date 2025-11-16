package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getGatewayResponseCmd = &cobra.Command{
	Use:   "get-gateway-response",
	Short: "Gets a GatewayResponse of a specified response type on the given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getGatewayResponseCmd).Standalone()

	apigateway_getGatewayResponseCmd.Flags().String("response-type", "", "The response type of the associated GatewayResponse.")
	apigateway_getGatewayResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_getGatewayResponseCmd.MarkFlagRequired("response-type")
	apigateway_getGatewayResponseCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_getGatewayResponseCmd)
}
