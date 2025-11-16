package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_updateDeploymentCmd = &cobra.Command{
	Use:   "update-deployment",
	Short: "Updates a Deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_updateDeploymentCmd).Standalone()

	apigatewayv2_updateDeploymentCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_updateDeploymentCmd.Flags().String("deployment-id", "", "The deployment ID.")
	apigatewayv2_updateDeploymentCmd.Flags().String("description", "", "The description for the deployment resource.")
	apigatewayv2_updateDeploymentCmd.MarkFlagRequired("api-id")
	apigatewayv2_updateDeploymentCmd.MarkFlagRequired("deployment-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_updateDeploymentCmd)
}
