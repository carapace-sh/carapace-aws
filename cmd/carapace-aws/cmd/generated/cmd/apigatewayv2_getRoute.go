package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getRouteCmd = &cobra.Command{
	Use:   "get-route",
	Short: "Gets a Route.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getRouteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_getRouteCmd).Standalone()

		apigatewayv2_getRouteCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_getRouteCmd.Flags().String("route-id", "", "The route ID.")
		apigatewayv2_getRouteCmd.MarkFlagRequired("api-id")
		apigatewayv2_getRouteCmd.MarkFlagRequired("route-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_getRouteCmd)
}
