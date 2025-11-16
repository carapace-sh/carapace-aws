package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getRouteResponseCmd = &cobra.Command{
	Use:   "get-route-response",
	Short: "Gets a RouteResponse.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getRouteResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getRouteResponseCmd).Standalone()

		apigatewayv2_getRouteResponseCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getRouteResponseCmd.Flags().String("route-id", "", "The route ID.")
		apigatewayv2_getRouteResponseCmd.Flags().String("route-response-id", "", "The route response ID.")
		apigatewayv2_getRouteResponseCmd.MarkFlagRequired("api-id")
		apigatewayv2_getRouteResponseCmd.MarkFlagRequired("route-id")
		apigatewayv2_getRouteResponseCmd.MarkFlagRequired("route-response-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getRouteResponseCmd)
}
