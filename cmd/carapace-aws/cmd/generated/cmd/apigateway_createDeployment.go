package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Creates a Deployment resource, which makes a specified RestApi callable over the internet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_createDeploymentCmd).Standalone()

	apigateway_createDeploymentCmd.Flags().String("cache-cluster-enabled", "", "Enables a cache cluster for the Stage resource specified in the input.")
	apigateway_createDeploymentCmd.Flags().String("cache-cluster-size", "", "The stage's cache capacity in GB.")
	apigateway_createDeploymentCmd.Flags().String("canary-settings", "", "The input configuration for the canary deployment when the deployment is a canary release deployment.")
	apigateway_createDeploymentCmd.Flags().String("description", "", "The description for the Deployment resource to create.")
	apigateway_createDeploymentCmd.Flags().String("rest-api-id", "", "The string identifier of the associated RestApi.")
	apigateway_createDeploymentCmd.Flags().String("stage-description", "", "The description of the Stage resource for the Deployment resource to create.")
	apigateway_createDeploymentCmd.Flags().String("stage-name", "", "The name of the Stage resource for the Deployment resource to create.")
	apigateway_createDeploymentCmd.Flags().String("tracing-enabled", "", "Specifies whether active tracing with X-ray is enabled for the Stage.")
	apigateway_createDeploymentCmd.Flags().String("variables", "", "A map that defines the stage variables for the Stage resource that is associated with the new deployment.")
	apigateway_createDeploymentCmd.MarkFlagRequired("rest-api-id")
	apigatewayCmd.AddCommand(apigateway_createDeploymentCmd)
}
