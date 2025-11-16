package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getGatewayResponsesCmd = &cobra.Command{
	Use:   "get-gateway-responses",
	Short: "Gets the GatewayResponses collection on the given RestApi.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getGatewayResponsesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getGatewayResponsesCmd).Standalone()

		apigateway_getGatewayResponsesCmd.Flags().String("limit", "", "The maximum number of returned results per page.")
		apigateway_getGatewayResponsesCmd.Flags().String("position", "", "The current pagination position in the paged result set.")
		apigateway_getGatewayResponsesCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getGatewayResponsesCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getGatewayResponsesCmd)
}
