package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateStageCmd = &cobra.Command{
	Use:   "update-stage",
	Short: "Updates a Stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateStageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_updateStageCmd).Standalone()

		apigatewayv2_updateStageCmd.Flags().String("access-log-settings", "", "Settings for logging access in this stage.")
		apigatewayv2_updateStageCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_updateStageCmd.Flags().String("auto-deploy", "", "Specifies whether updates to an API automatically trigger a new deployment.")
		apigatewayv2_updateStageCmd.Flags().String("client-certificate-id", "", "The identifier of a client certificate for a Stage.")
		apigatewayv2_updateStageCmd.Flags().String("default-route-settings", "", "The default route settings for the stage.")
		apigatewayv2_updateStageCmd.Flags().String("deployment-id", "", "The deployment identifier for the API stage.")
		apigatewayv2_updateStageCmd.Flags().String("description", "", "The description for the API stage.")
		apigatewayv2_updateStageCmd.Flags().String("route-settings", "", "Route settings for the stage.")
		apigatewayv2_updateStageCmd.Flags().String("stage-name", "", "The stage name.")
		apigatewayv2_updateStageCmd.Flags().String("stage-variables", "", "A map that defines the stage variables for a Stage.")
		apigatewayv2_updateStageCmd.MarkFlagRequired("api-id")
		apigatewayv2_updateStageCmd.MarkFlagRequired("stage-name")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateStageCmd)
}
