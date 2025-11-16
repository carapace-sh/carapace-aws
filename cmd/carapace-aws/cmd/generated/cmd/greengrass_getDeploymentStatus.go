package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getDeploymentStatusCmd = &cobra.Command{
	Use:   "get-deployment-status",
	Short: "Returns the status of a deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getDeploymentStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_getDeploymentStatusCmd).Standalone()

		greengrass_getDeploymentStatusCmd.Flags().String("deployment-id", "", "The ID of the deployment.")
		greengrass_getDeploymentStatusCmd.Flags().String("group-id", "", "The ID of the Greengrass group.")
		greengrass_getDeploymentStatusCmd.MarkFlagRequired("deployment-id")
		greengrass_getDeploymentStatusCmd.MarkFlagRequired("group-id")
	})
	greengrassCmd.AddCommand(greengrass_getDeploymentStatusCmd)
}
