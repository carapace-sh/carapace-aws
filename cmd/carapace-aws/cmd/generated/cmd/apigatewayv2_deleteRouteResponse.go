package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteRouteResponseCmd = &cobra.Command{
	Use:   "delete-route-response",
	Short: "Deletes a RouteResponse.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteRouteResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_deleteRouteResponseCmd).Standalone()

		apigatewayv2_deleteRouteResponseCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_deleteRouteResponseCmd.Flags().String("route-id", "", "The route ID.")
		apigatewayv2_deleteRouteResponseCmd.Flags().String("route-response-id", "", "The route response ID.")
		apigatewayv2_deleteRouteResponseCmd.MarkFlagRequired("api-id")
		apigatewayv2_deleteRouteResponseCmd.MarkFlagRequired("route-id")
		apigatewayv2_deleteRouteResponseCmd.MarkFlagRequired("route-response-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteRouteResponseCmd)
}
