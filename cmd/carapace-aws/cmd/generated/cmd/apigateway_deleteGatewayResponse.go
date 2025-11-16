package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteGatewayResponseCmd = &cobra.Command{
	Use:   "delete-gateway-response",
	Short: "Clears any customization of a GatewayResponse of a specified response type on the given RestApi and resets it with the default settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteGatewayResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteGatewayResponseCmd).Standalone()

		apigateway_deleteGatewayResponseCmd.Flags().String("response-type", "", "The response type of the associated GatewayResponse.")
		apigateway_deleteGatewayResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_deleteGatewayResponseCmd.MarkFlagRequired("response-type")
		apigateway_deleteGatewayResponseCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteGatewayResponseCmd)
}
