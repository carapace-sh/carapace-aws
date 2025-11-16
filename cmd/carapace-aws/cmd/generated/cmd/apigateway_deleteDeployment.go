package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_deleteDeploymentCmd = &cobra.Command{
	Use:   "delete-deployment",
	Short: "Deletes a Deployment resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_deleteDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_deleteDeploymentCmd).Standalone()

		apigateway_deleteDeploymentCmd.Flags().String("deployment-id", "", "The identifier of the Deployment resource to delete.")
		apigateway_deleteDeploymentCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_deleteDeploymentCmd.MarkFlagRequired("deployment-id")
		apigateway_deleteDeploymentCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_deleteDeploymentCmd)
}
