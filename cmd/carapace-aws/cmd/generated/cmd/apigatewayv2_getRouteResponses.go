package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getRouteResponsesCmd = &cobra.Command{
	Use:   "get-route-responses",
	Short: "Gets the RouteResponses for a Route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getRouteResponsesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getRouteResponsesCmd).Standalone()

		apigatewayv2_getRouteResponsesCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getRouteResponsesCmd.Flags().String("max-results", "", "The maximum number of elements to be returned for this resource.")
		apigatewayv2_getRouteResponsesCmd.Flags().String("next-token", "", "The next page of elements from this collection.")
		apigatewayv2_getRouteResponsesCmd.Flags().String("route-id", "", "The route ID.")
		apigatewayv2_getRouteResponsesCmd.MarkFlagRequired("api-id")
		apigatewayv2_getRouteResponsesCmd.MarkFlagRequired("route-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getRouteResponsesCmd)
}
