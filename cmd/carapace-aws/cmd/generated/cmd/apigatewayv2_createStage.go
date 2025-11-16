package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createStageCmd = &cobra.Command{
	Use:   "create-stage",
	Short: "Creates a Stage for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_createStageCmd).Standalone()

		apigatewayv2_createStageCmd.Flags().String("access-log-settings", "", "Settings for logging access in this stage.")
		apigatewayv2_createStageCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_createStageCmd.Flags().String("auto-deploy", "", "Specifies whether updates to an API automatically trigger a new deployment.")
		apigatewayv2_createStageCmd.Flags().String("client-certificate-id", "", "The identifier of a client certificate for a Stage.")
		apigatewayv2_createStageCmd.Flags().String("default-route-settings", "", "The default route settings for the stage.")
		apigatewayv2_createStageCmd.Flags().String("deployment-id", "", "The deployment identifier of the API stage.")
		apigatewayv2_createStageCmd.Flags().String("description", "", "The description for the API stage.")
		apigatewayv2_createStageCmd.Flags().String("route-settings", "", "Route settings for the stage, by routeKey.")
		apigatewayv2_createStageCmd.Flags().String("stage-name", "", "The name of the stage.")
		apigatewayv2_createStageCmd.Flags().String("stage-variables", "", "A map that defines the stage variables for a Stage.")
		apigatewayv2_createStageCmd.Flags().String("tags", "", "The collection of tags.")
		apigatewayv2_createStageCmd.MarkFlagRequired("api-id")
		apigatewayv2_createStageCmd.MarkFlagRequired("stage-name")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_createStageCmd)
}
