package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getDeploymentCmd = &cobra.Command{
	Use:   "get-deployment",
	Short: "Gets information about a Deployment resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_getDeploymentCmd).Standalone()

		apigateway_getDeploymentCmd.Flags().String("deployment-id", "", "The identifier of the Deployment resource to get information about.")
		apigateway_getDeploymentCmd.Flags().String("embed", "", "A query parameter to retrieve the specified embedded resources of the returned Deployment resource in the response.")
		apigateway_getDeploymentCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
		apigateway_getDeploymentCmd.MarkFlagRequired("deployment-id")
		apigateway_getDeploymentCmd.MarkFlagRequired("rest-api-id")
	})
	apigatewayCmd.AddCommand(apigateway_getDeploymentCmd)
}
