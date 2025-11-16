package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createDeploymentCmd = &cobra.Command{
	Use:   "create-deployment",
	Short: "Creates a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createDeploymentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createDeploymentCmd).Standalone()

		greengrass_createDeploymentCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createDeploymentCmd.Flags().String("deployment-id", "", "The ID of the deployment if you wish to redeploy a previous deployment.")
		greengrass_createDeploymentCmd.Flags().String("deployment-type", "", "The type of deployment.")
		greengrass_createDeploymentCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_createDeploymentCmd.Flags().String("group-version-id", "", "The ID of the group version to be deployed.")
		greengrass_createDeploymentCmd.MarkFlagRequired("deployment-type")
		greengrass_createDeploymentCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_createDeploymentCmd)
}
