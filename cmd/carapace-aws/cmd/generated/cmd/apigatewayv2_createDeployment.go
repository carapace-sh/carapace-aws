package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Creates a Deployment for an API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_createDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_createDeploymentCmd).Standalone()

		apigatewayv2_createDeploymentCmd.Flags().String("api-id", "", "The API identifier.")
		apigatewayv2_createDeploymentCmd.Flags().String("description", "", "The description for the deployment resource.")
		apigatewayv2_createDeploymentCmd.Flags().String("stage-name", "", "The name of the Stage resource for the Deployment resource to create.")
		apigatewayv2_createDeploymentCmd.MarkFlagRequired("api-id")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_createDeploymentCmd)
}
