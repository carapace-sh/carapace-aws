package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_updateDeploymentCmd = &cobra.Command{
	Use:   "update-deployment",
	Short: "Changes information about a Deployment resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_updateDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_updateDeploymentCmd).Standalone()

		apigateway_updateDeploymentCmd.Flags().String("deployment-id", "", "The replacement identifier for the Deployment resource to change information about.")
		apigateway_updateDeploymentCmd.Flags().String("patch-operations", "", "For more information about supported patch operations, see [Patch Operations](https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).")
		apigateway_updateDeploymentCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_updateDeploymentCmd.MarkFlagRequired("deployment-id")
		apigateway_updateDeploymentCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_updateDeploymentCmd)
}
