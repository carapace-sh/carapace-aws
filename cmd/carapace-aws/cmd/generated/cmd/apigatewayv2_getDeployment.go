package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Gets a Deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getDeploymentCmd).Standalone()

	apigatewayv2_getDeploymentCmd.Flags().String("api-id", "", "The API identifier.")
	apigatewayv2_getDeploymentCmd.Flags().String("deployment-id", "", "The deployment ID.")
	apigatewayv2_getDeploymentCmd.MarkFlagRequired("api-id")
	apigatewayv2_getDeploymentCmd.MarkFlagRequired("deployment-id")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getDeploymentCmd)
}
