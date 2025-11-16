package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_deleteDeploymentCmd = &cobra.Command{
	Use:   "delete-deployment",
	Short: "Deletes a Deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_deleteDeploymentCmd).Standalone()

	apigatewayv2_deleteDeploymentCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_deleteDeploymentCmd.Flags().String("deployment-id", "", "The deployment ID.")
	apigatewayv2_deleteDeploymentCmd.MarkFlagRequired("api-id")
	apigatewayv2_deleteDeploymentCmd.MarkFlagRequired("deployment-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_deleteDeploymentCmd)
}
