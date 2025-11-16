package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateGatewayResponseCmd = &cobra.Command{
	Use:   "update-gateway-response",
	Short: "Updates a GatewayResponse of a specified response type on the given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateGatewayResponseCmd).Standalone()

	apigateway_updateGatewayResponseCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
	apigateway_updateGatewayResponseCmd.Flags().String("response-type", "", "The response type of the associated GatewayResponse.")
	apigateway_updateGatewayResponseCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_updateGatewayResponseCmd.MarkFlagRequired("response-type")
	apigateway_updateGatewayResponseCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_updateGatewayResponseCmd)
}
