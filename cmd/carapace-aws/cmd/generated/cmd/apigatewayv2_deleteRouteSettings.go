package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteRouteSettingsCmd = &cobra.Command{
	Use:   "delete-route-settings",
	Short: "Deletes the RouteSettings for a stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteRouteSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_deleteRouteSettingsCmd).Standalone()

		apigatewayv2_deleteRouteSettingsCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_deleteRouteSettingsCmd.Flags().String("route-key", "", "The route key.")
		apigatewayv2_deleteRouteSettingsCmd.Flags().String("stage-name", "", "The stage name.")
		apigatewayv2_deleteRouteSettingsCmd.MarkFlagRequired("api-id")
		apigatewayv2_deleteRouteSettingsCmd.MarkFlagRequired("route-key")
		apigatewayv2_deleteRouteSettingsCmd.MarkFlagRequired("stage-name")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteRouteSettingsCmd)
}
